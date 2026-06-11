<script lang="ts">
  import { onMount, onDestroy } from "svelte";

  let { data }: { data: any } = $props();

  interface ChatRoom {
    id: string;
    title: string;
    project_id: string;
    lastMessage?: string;
    lastMessageTime?: string;
    unread?: number;
    [key: string]: any;
  }

  interface ChatMessage {
    id: string;
    content: string;
    sender_id: string;
    group_id?: string;
    type: string;
    file_url?: string;
    file_name?: string;
    created_at: string;
    _id?: string;
    user_name?: string;
    avatar?: string;
  }

  interface TeamMember {
    id: string;
    name: string;
    avatar: string;
    status: string;
    title?: string;
    role?: string;
    isCurrentUser?: boolean;
  }

  interface SharedFile {
    url: string;
    name: string;
    sender_id: string;
    created_at: string;
  }

  import { PUBLIC_API_URL, PUBLIC_WS_URL } from '$env/static/public';

  let chatRooms: ChatRoom[] = $state([]);
  let messages: ChatMessage[] = $state([]);
  let teamMembers: TeamMember[] = $state([]);
  let sharedFiles: SharedFile[] = $state([]);

  let selectedRoom: ChatRoom | null = $state(null);
  let messageInput: string = $state("");
  let searchQuery: string = $state("");
  let filteredRooms: ChatRoom[] = $state([]);
  let ws: WebSocket | null = $state(null);
  let isUploading: boolean = $state(false);
  let uploadProgress: number = $state(0);

  let currentUserId: string = $state("");
  let currentUserName: string = $state("");
  let currentUserAvatar: string = $state("");
  let chatWsReconnectTimer: any = null;
  let shouldReconnectChatWs = true;

  onMount(async () => {
    currentUserId = data?.id || "";
    currentUserName = data?.userName || "";
    currentUserAvatar = data?.avatarUrl || `${PUBLIC_API_URL}/avatars/${data?.id}`;

    try {
      const res = await fetch(`/api/chats`);
      if (res.ok) {
        const rooms = await res.json() || [];
        chatRooms = rooms.map((r: any) => ({
          ...r,
          id: r.project_id || r.id,
          title: r.title,
        }));
        filteredRooms = chatRooms;
        if (chatRooms.length > 0) {
          selectRoom(chatRooms[0]);
        }
      }
    } catch (e) {
      console.error("Failed to fetch chat rooms", e);
    }
  });

  onDestroy(() => {
    shouldReconnectChatWs = false;
    if (chatWsReconnectTimer) clearTimeout(chatWsReconnectTimer);
    if (ws) {
      ws.close();
    }
  });

  async function loadRoomDetails(roomId: string) {
    try {
      const res = await fetch(`/api/chats/${roomId}`);
      if (res.ok) {
        const roomData = await res.json();
        const rawMessages: ChatMessage[] = roomData.messages || [];

        const usersMap = new Map<string, TeamMember>();
        if (roomData.members) {
          roomData.members.forEach((m: any) => {
            usersMap.set(m._id || m.id, {
              id: m._id || m.id,
              name: m.name || "Unknown",
              avatar: `${PUBLIC_API_URL}/avatars/${m._id || m.id}`,
              status: "online",
              title: m.title || "",
              isCurrentUser: (m._id || m.id) === currentUserId
            });
          });
          teamMembers = Array.from(usersMap.values());
        }

        messages = rawMessages.map((msg) => {
          const sender = usersMap.get(msg.sender_id);
          return {
            ...msg,
            user_name: sender?.name || "User",
            avatar: sender?.avatar || `${PUBLIC_API_URL}/avatars/${msg.sender_id}`,
            file_url: msg.file_url ? (msg.file_url.startsWith("http") ? msg.file_url : `${PUBLIC_API_URL}${msg.file_url}`) : "",
            created_at: msg.created_at || new Date().toISOString(),
          };
        });

        const files: SharedFile[] = rawMessages
          .filter((m: ChatMessage) => m.type === "file" || m.file_url)
          .map((m: ChatMessage) => ({
            url: m.file_url ? (m.file_url.startsWith("http") ? m.file_url : `${PUBLIC_API_URL}${m.file_url}`) : "",
            name: m.file_name || "file",
            sender_id: m.sender_id,
            created_at: m.created_at,
          }));
        sharedFiles = files;

        setTimeout(() => {
          const container = document.getElementById("chat-messages");
          if (container) container.scrollTop = container.scrollHeight;
        }, 50);
      }
    } catch(e) {
      console.error("Failed to load room details", e);
    }
  }

  $effect(() => {
    if (searchQuery) {
      filteredRooms = chatRooms.filter((room) =>
        room.title.toLowerCase().includes(searchQuery.toLowerCase())
      );
    } else {
      filteredRooms = chatRooms;
    }
  });

  async function selectRoom(room: ChatRoom) {
    if (chatWsReconnectTimer) clearTimeout(chatWsReconnectTimer);
    selectedRoom = room;
    await loadRoomDetails(room.id);
    connectChatWs(room.id);
    setTimeout(() => {
      const container = document.getElementById("chat-messages");
      if (container) container.scrollTop = container.scrollHeight;
    }, 50);
  }

  function connectChatWs(projectId: string, retryCount = 0) {
    if (ws) {
      ws.close();
    }
    const userId = currentUserId;
    if (!userId) return;

    ws = new WebSocket(`${PUBLIC_WS_URL}/ws/chats/${projectId}/${userId}`);

    ws.onmessage = (event) => {
      try {
        const msg = JSON.parse(event.data);
        if (msg && (msg._id || msg.id)) {
          const sender = teamMembers.find(m => m.id === msg.sender_id);

          const fileUrl = msg.file_url ? (msg.file_url.startsWith("http") ? msg.file_url : `${PUBLIC_API_URL}${msg.file_url}`) : "";

          if (msg.sender_id === currentUserId && msg.client_id) {
            const optIndex = messages.findIndex(m => m.id === msg.client_id);
            if (optIndex !== -1) {
              const newMessages = [...messages];
              newMessages[optIndex] = {
                id: msg._id || msg.id,
                content: msg.content || "",
                sender_id: msg.sender_id,
                group_id: msg.group_id,
                type: msg.type || "text",
                file_url: fileUrl,
                file_name: msg.file_name || "",
                created_at: msg.created_at || new Date().toISOString(),
                user_name: sender?.name || msg.user_name || "User",
                avatar: sender?.avatar || `${PUBLIC_API_URL}/avatars/${msg.sender_id}`,
              };
              messages = newMessages;
              return;
            }
          }

          messages = [...messages, {
            id: msg._id || msg.id,
            content: msg.content || "",
            sender_id: msg.sender_id,
            group_id: msg.group_id,
            type: msg.type || "text",
            file_url: fileUrl,
            file_name: msg.file_name || "",
            created_at: msg.created_at || new Date().toISOString(),
            user_name: sender?.name || msg.user_name || "User",
            avatar: sender?.avatar || `${PUBLIC_API_URL}/avatars/${msg.sender_id}`,
          }];

          setTimeout(() => {
            const container = document.getElementById("chat-messages");
            if (container) container.scrollTop = container.scrollHeight;
          }, 50);
        }
      } catch (e) {
        console.error("Error parsing message", e);
      }
    };

    ws.onclose = () => {
      if (!shouldReconnectChatWs) return;
      const delay = Math.min(1000 * Math.pow(2, retryCount), 30000);
      chatWsReconnectTimer = setTimeout(() => {
        connectChatWs(projectId, retryCount + 1);
      }, delay);
    };

    ws.onerror = () => {
      ws?.close();
    };
  }

  function sendMessage() {
    if (!messageInput.trim()) return;

    if (ws && ws.readyState === WebSocket.OPEN) {
      const optimisticId = "opt_" + Date.now();
      const optMsg = {
        id: optimisticId,
        content: messageInput.trim(),
        sender_id: currentUserId,
        type: "text",
        created_at: new Date().toISOString(),
        user_name: currentUserName,
        avatar: currentUserAvatar,
      };
      messages = [...messages, optMsg];

      ws.send(JSON.stringify({ content: messageInput.trim(), type: "text", client_id: optimisticId }));

      setTimeout(() => {
        const container = document.getElementById("chat-messages");
        if (container) container.scrollTop = container.scrollHeight;
      }, 50);
    }

    messageInput = "";
  }

  function handleFileUpload() {
    const input = document.createElement("input");
    input.type = "file";
    input.accept = "image/*,.pdf,.doc,.docx,.txt,.zip,.rar,.csv,.xlsx,.xls";
    input.onchange = async () => {
      const file = input.files?.[0];
      if (!file) return;
      if (file.size > 10 * 1024 * 1024) {
        alert("File too large (max 10MB)");
        return;
      }

      isUploading = true;
      uploadProgress = 0;

      try {
        const formData = new FormData();
        formData.append("file", file);

        const xhr = new XMLHttpRequest();
        xhr.open("POST", `/api/chats/upload`, true);

        xhr.upload.onprogress = (e) => {
          if (e.lengthComputable) {
            uploadProgress = Math.round((e.loaded / e.total) * 100);
          }
        };

        const result = await new Promise<any>((resolve, reject) => {
          xhr.onload = () => {
            if (xhr.status === 200) resolve(JSON.parse(xhr.responseText));
            else reject(new Error("Upload failed"));
          };
          xhr.onerror = () => reject(new Error("Upload failed"));
          xhr.send(formData);
        });

        if (ws && ws.readyState === WebSocket.OPEN) {
          const optimisticId = "opt_file_" + Date.now();
          messages = [...messages, {
            id: optimisticId,
            content: file.name,
            sender_id: currentUserId,
            type: "file",
            file_url: result.url,
            file_name: result.name,
            created_at: new Date().toISOString(),
            user_name: currentUserName,
            avatar: currentUserAvatar,
          }];

          ws.send(JSON.stringify({
            content: file.name,
            type: "file",
            file_url: result.url,
            file_name: result.name,
            client_id: optimisticId,
          }));

          setTimeout(() => {
            const container = document.getElementById("chat-messages");
            if (container) container.scrollTop = container.scrollHeight;
          }, 50);
        }
      } catch (e) {
        console.error("Upload error", e);
        alert("Failed to upload file");
      } finally {
        isUploading = false;
        uploadProgress = 0;
      }
    };
    input.click();
  }

  function formatTime(timestamp: string) {
    if (!timestamp) return "";
    const date = new Date(timestamp);
    return date.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
  }

  function formatDate(timestamp: string) {
    if (!timestamp) return "";
    const date = new Date(timestamp);
    const today = new Date();
    const yesterday = new Date(today);
    yesterday.setDate(yesterday.getDate() - 1);

    if (date.toDateString() === today.toDateString()) {
      return "Today";
    } else if (date.toDateString() === yesterday.toDateString()) {
      return "Yesterday";
    } else {
      return date.toLocaleDateString("en-US", {
        month: "short",
        day: "numeric",
      });
    }
  }

  function handleKeyDown(event: KeyboardEvent) {
    if (event.key === "Enter" && !event.shiftKey) {
      event.preventDefault();
      sendMessage();
    }
  }

  function isImage(url: string): boolean {
    return /\.(jpg|jpeg|png|gif|webp)$/i.test(url);
  }

  function getFileIcon(name: string): string {
    const ext = name.split('.').pop()?.toLowerCase() || "";
    if (["pdf"].includes(ext)) return "pdf";
    if (["doc", "docx"].includes(ext)) return "doc";
    if (["xls", "xlsx", "csv"].includes(ext)) return "sheet";
    if (["zip", "rar", "7z"].includes(ext)) return "zip";
    if (["txt"].includes(ext)) return "txt";
    return "generic";
  }

  function formatFileSize(bytes: number): string {
    if (bytes < 1024) return bytes + " B";
    if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + " KB";
    return (bytes / (1024 * 1024)).toFixed(1) + " MB";
  }

  function isCurrentUser(senderId: string): boolean {
    return senderId === currentUserId;
  }
</script>

<main class="flex-1 flex flex-col p-4 h-screen bg-gray-50">
  <div class="flex gap-4 flex-1 h-full overflow-hidden">
    <!-- Chat Rooms List -->
    <div class="bg-white rounded-xl shadow-sm p-3 w-64 flex flex-col flex-shrink-0 border border-gray-100">
      <div class="flex items-center justify-between mb-3">
        <h3 class="text-sm font-semibold text-gray-800">Conversations</h3>
        <span class="text-[10px] text-gray-400 bg-gray-100 px-1.5 py-0.5 rounded-full">{chatRooms.length}</span>
      </div>

      <div class="relative mb-2">
        <input
          type="text"
          placeholder="Search conversations..."
          bind:value={searchQuery}
          class="w-full p-1.5 pl-7 text-xs border border-gray-200 rounded-lg focus:outline-none focus:ring-1 focus:ring-[#6b48ff] bg-gray-50"
        />
        <svg class="w-3.5 h-3.5 absolute left-2 top-2 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </div>

      <div class="flex-1 overflow-y-auto custom-scrollbar -mx-1">
        {#if filteredRooms.length === 0}
          <div class="text-center text-gray-400 text-xs py-8">No conversations found</div>
        {/if}
        <div class="space-y-0.5 px-1">
          {#each filteredRooms as room}
            <div
              class={`p-2 rounded-lg cursor-pointer hover:bg-gray-50 flex items-center transition-colors ${
                selectedRoom?.id === room.id ? "bg-[#6b48ff] bg-opacity-10 border-l-2 border-[#6b48ff]" : "border-l-2 border-transparent"
              }`}
              role="button" tabindex="0"
              onclick={() => selectRoom(room)}
              onkeydown={(e) => { if (e.key === 'Enter') selectRoom(room); }}
            >
              <div class="w-8 h-8 rounded-full bg-gradient-to-br from-purple-500 to-indigo-600 flex items-center justify-center flex-shrink-0 mr-2">
                <span class="text-white font-medium text-xs">{room.title.substring(0, 2).toUpperCase()}</span>
              </div>
              <div class="flex-1 overflow-hidden min-w-0">
                <div class="flex justify-between items-center">
                  <span class="text-xs font-medium text-gray-800 truncate">{room.title}</span>
                  <span class="text-[10px] text-gray-400 ml-1 flex-shrink-0">{room.lastMessageTime || ""}</span>
                </div>
                <div class="flex justify-between items-center mt-0.5">
                  <p class="text-[10px] text-gray-500 truncate">{room.lastMessage || "No messages yet"}</p>
                  {#if room.unread && room.unread > 0}
                    <span class="flex-shrink-0 bg-[#6b48ff] text-white text-[10px] font-medium rounded-full min-w-[16px] h-4 flex items-center justify-center px-1 ml-1">{room.unread}</span>
                  {/if}
                </div>
              </div>
            </div>
          {/each}
        </div>
      </div>
    </div>

    <!-- Chat Window -->
    <div class="bg-white rounded-xl shadow-sm flex-1 flex flex-col overflow-hidden border border-gray-100">
      {#if selectedRoom}
        <!-- Chat Header -->
        <div class="flex justify-between items-center px-4 py-3 border-b border-gray-100 bg-white">
          <div class="flex items-center">
            <div class="w-9 h-9 rounded-full bg-gradient-to-br from-purple-500 to-indigo-600 flex items-center justify-center mr-3">
              <span class="text-white font-semibold text-sm">{selectedRoom.title.substring(0, 2).toUpperCase()}</span>
            </div>
            <div>
              <h3 class="text-sm font-semibold text-gray-800">{selectedRoom.title}</h3>
              <div class="flex items-center gap-1">
                <span class="w-1.5 h-1.5 rounded-full bg-green-500"></span>
                <p class="text-[10px] text-gray-500">
                  {teamMembers.filter((m) => m.status === "online").length} online
                </p>
              </div>
            </div>
          </div>
          <div class="flex items-center space-x-1">
            <button class="p-1.5 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors" title="Search">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </button>
            <button class="p-1.5 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors" title="More">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Messages -->
        <div id="chat-messages" class="flex-1 overflow-y-auto p-4 space-y-3 bg-gradient-to-b from-gray-50 to-white custom-scrollbar">
          {#each messages as message, i}
            {#if i === 0 || formatDate(messages[i - 1]?.created_at) !== formatDate(message.created_at)}
              <div class="flex justify-center my-2">
                <div class="bg-white text-gray-500 text-[10px] font-medium px-3 py-0.5 rounded-full shadow-sm border border-gray-100">
                  {formatDate(message.created_at)}
                </div>
              </div>
            {/if}

            <div class={`flex ${isCurrentUser(message.sender_id) ? "justify-end" : "justify-start"} items-end gap-2`}>
              {#if !isCurrentUser(message.sender_id)}
                <img
                  src={message.avatar || `${PUBLIC_API_URL}/avatars/${message.sender_id}`}
                  alt={message.user_name}
                  class="w-7 h-7 rounded-full flex-shrink-0 border border-gray-200 object-cover"
                />
              {/if}

              <div class={`max-w-[70%] ${isCurrentUser(message.sender_id) ? "items-end" : "items-start"} flex flex-col`}>
                {#if !isCurrentUser(message.sender_id)}
                  <span class="text-[10px] font-semibold text-gray-600 mb-0.5 ml-1">{message.user_name}</span>
                {/if}

                {#if message.type === "file"}
                  <div class={`rounded-2xl overflow-hidden ${
                    isCurrentUser(message.sender_id)
                      ? "bg-[#6b48ff] rounded-tr-sm"
                      : "bg-white border border-gray-100 shadow-sm rounded-tl-sm"
                  }`}>
                    {#if message.file_url && isImage(message.file_url)}
                      <div class="max-w-[280px]">
                        <button
                          type="button"
                          class="w-full p-0 border-0 cursor-pointer block"
                          onclick={() => window.open(message.file_url, '_blank')}
                        >
                          <img
                            src={message.file_url}
                            alt={message.file_name || "Image"}
                            class="w-full h-auto object-cover max-h-[300px]"
                          />
                        </button>
                        <div class={`px-3 py-1.5 flex items-center justify-between text-xs ${
                          isCurrentUser(message.sender_id) ? "text-purple-200" : "text-gray-500"
                        }`}>
                          <span class="truncate">{message.file_name}</span>
                          <a href={message.file_url} download class="ml-2 hover:opacity-80" title="Download" aria-label="Download file">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
                            </svg>
                          </a>
                        </div>
                      </div>
                    {:else}
                      <div class="flex items-center gap-2 p-2.5 min-w-[200px]">
                        <div class={`p-2 rounded-lg ${
                          isCurrentUser(message.sender_id) ? "bg-purple-500" : "bg-gray-100"
                        }`}>
                          {#if getFileIcon(message.file_name || "") === "pdf"}
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-red-500" viewBox="0 0 20 20" fill="currentColor">
                              <path fill-rule="evenodd" d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" clip-rule="evenodd"/>
                            </svg>
                          {:else}
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                            </svg>
                          {/if}
                        </div>
                        <div class="flex-1 min-w-0">
                          <p class={`text-xs font-medium truncate ${
                            isCurrentUser(message.sender_id) ? "text-white" : "text-gray-800"
                          }`}>{message.file_name}</p>
                        </div>
                        <a
                          href={message.file_url}
                          download
                          class={`p-1.5 rounded-full transition-colors flex-shrink-0 ${
                            isCurrentUser(message.sender_id)
                              ? "hover:bg-purple-500 text-white"
                              : "hover:bg-gray-100 text-gray-500"
                          }`}
                            title="Download"
                            aria-label="Download file"
                          >
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
                            </svg>
                          </a>
                      </div>
                    {/if}
                  </div>
                {:else}
                  <div class={`px-3 py-2 text-xs leading-relaxed ${
                    isCurrentUser(message.sender_id)
                      ? "bg-[#6b48ff] text-white rounded-2xl rounded-tr-sm"
                      : "bg-white border border-gray-100 shadow-sm text-gray-800 rounded-2xl rounded-tl-sm"
                  }`}>
                    <p>{message.content}</p>
                  </div>
                {/if}

                <span class={`text-[9px] mt-0.5 px-1 ${
                  isCurrentUser(message.sender_id) ? "text-gray-400" : "text-gray-400"
                }`}>
                  {formatTime(message.created_at)}
                  {#if message.id && message.id.startsWith("opt_")}
                    <span class="ml-1">Sending...</span>
                  {/if}
                </span>
              </div>

              {#if isCurrentUser(message.sender_id)}
                <img
                  src={message.avatar || data?.avatarUrl || `${PUBLIC_API_URL}/avatars/${currentUserId}`}
                  alt={currentUserName}
                  class="w-7 h-7 rounded-full flex-shrink-0 border border-gray-200 object-cover"
                />
              {/if}
            </div>
          {/each}

          {#if messages.length === 0}
            <div class="flex flex-col items-center justify-center h-full text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
              </svg>
              <p class="text-xs">No messages yet. Start a conversation!</p>
            </div>
          {/if}
        </div>

        <!-- Input -->
        <div class="p-3 border-t border-gray-100 bg-white">
          {#if isUploading}
            <div class="mb-2 flex items-center gap-2">
              <div class="flex-1 bg-gray-200 rounded-full h-1.5">
                <div class="bg-[#6b48ff] h-1.5 rounded-full transition-all" style="width: {uploadProgress}%"></div>
              </div>
              <span class="text-[10px] text-gray-500">{uploadProgress}%</span>
            </div>
          {/if}
          <div class="flex gap-2 items-center bg-gray-50 p-1.5 rounded-xl border border-gray-200 focus-within:border-[#6b48ff] focus-within:ring-1 focus-within:ring-[#6b48ff] transition-all">
            <button
              class="p-1.5 rounded-lg hover:bg-gray-200 text-gray-400 hover:text-gray-600 transition-colors disabled:opacity-50"
              onclick={handleFileUpload}
              disabled={isUploading}
              title="Attach file"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M8 4a3 3 0 00-3 3v4a5 5 0 0010 0V7a1 1 0 112 0v4a7 7 0 11-14 0V7a5 5 0 0110 0v4a3 3 0 11-6 0V7a1 1 0 012 0v4a1 1 0 102 0V7a3 3 0 00-3-3z" clip-rule="evenodd" />
              </svg>
            </button>

            <input
              type="text"
              bind:value={messageInput}
              onkeydown={handleKeyDown}
              class="flex-1 p-1 text-xs bg-transparent border-none focus:outline-none placeholder-gray-400"
              placeholder="Type your message..."
            />

            <button
              class="bg-[#6b48ff] text-white p-1.5 rounded-lg transition-all hover:bg-[#5a3de6] disabled:bg-gray-300 disabled:cursor-not-allowed"
              onclick={sendMessage}
              disabled={!messageInput.trim() || isUploading}
              aria-label="Send message"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
                <path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z" />
              </svg>
            </button>
          </div>
        </div>
      {:else}
        <div class="flex-1 flex items-center justify-center text-gray-400">
          <div class="text-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
            </svg>
            <p class="text-sm font-medium">Select a conversation</p>
            <p class="text-xs mt-1">Choose a chat room to start messaging</p>
          </div>
        </div>
      {/if}
    </div>

    <!-- Team Members & Files Panel -->
    <div class="bg-white rounded-xl shadow-sm p-3 w-64 flex flex-col flex-shrink-0 border border-gray-100">
      {#if selectedRoom}
        <div class="mb-3">
          <h3 class="text-xs font-semibold text-gray-800 mb-2 flex items-center">
            Team Members
            <span class="ml-1.5 text-[10px] bg-[#6b48ff] text-white px-1.5 py-0.5 rounded-full">{teamMembers.length}</span>
          </h3>
          <div class="space-y-1 overflow-y-auto custom-scrollbar max-h-[35vh]">
            {#each teamMembers as member}
              <div class="flex items-center p-1.5 hover:bg-gray-50 rounded-lg transition-colors cursor-pointer">
                <div class="relative flex-shrink-0">
                  <img
                    src={member.avatar}
                    alt={member.name}
                    class="w-7 h-7 rounded-full object-cover border border-gray-200"
                  />
                  <div class={`absolute -bottom-0.5 -right-0.5 w-2.5 h-2.5 rounded-full border-2 border-white ${
                    member.status === "online" ? "bg-green-500" : "bg-gray-300"
                  }`}></div>
                </div>
                <div class="ml-2 flex-1 min-w-0">
                  <div class="flex items-center gap-1">
                    <p class="text-xs font-medium text-gray-800 truncate">{member.name}</p>
                    {#if member.isCurrentUser}
                      <span class="text-[8px] bg-blue-100 text-blue-700 px-1 rounded-full font-medium">You</span>
                    {/if}
                  </div>
                  <p class="text-[9px] text-gray-500 truncate">{member.title || member.role || ""}</p>
                </div>
              </div>
            {/each}
          </div>
        </div>

        <div class="border-t border-gray-100 pt-3">
          <h3 class="text-xs font-semibold text-gray-800 mb-2 flex items-center justify-between">
            <span>Shared Files</span>
            <span class="text-[10px] text-gray-400">{sharedFiles.length}</span>
          </h3>
          <div class="space-y-1.5 overflow-y-auto custom-scrollbar max-h-[40vh]">
            {#each sharedFiles as file}
              <a href={file.url} download target="_blank" class="block p-2 bg-gray-50 rounded-lg border border-gray-100 hover:border-gray-200 hover:bg-gray-100 transition-colors">
                <div class="flex items-center gap-2">
                  <div class="p-1.5 rounded-lg bg-blue-100 text-blue-600 flex-shrink-0">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" clip-rule="evenodd"/>
                    </svg>
                  </div>
                  <div class="flex-1 min-w-0">
                    <p class="text-[10px] font-medium text-gray-700 truncate">{file.name}</p>
                    <p class="text-[8px] text-gray-400">{formatTime(file.created_at)}</p>
                  </div>
                </div>
              </a>
            {/each}
            {#if sharedFiles.length === 0}
              <p class="text-[10px] text-gray-400 text-center py-4">No files shared yet</p>
            {/if}
          </div>
        </div>
      {:else}
        <div class="flex flex-col items-center justify-center h-full text-gray-400">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
          </svg>
          <p class="text-xs font-medium">No team selected</p>
        </div>
      {/if}
    </div>
  </div>
</main>
