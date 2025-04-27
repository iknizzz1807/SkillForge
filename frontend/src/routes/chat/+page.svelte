<script lang="ts">
  import { onMount, onDestroy } from "svelte";
  import { browser } from "$app/environment";
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  // State variables
  let chatRooms = $state<any[]>([]);
  let selectedRoom = $state<any>(null);
  let messages = $state<any[]>([]);
  let messageInput = $state("");
  let searchQuery = $state("");
  let socket: WebSocket | null = $state(null);
  let teamMembers = $state<any[]>([]);
  let sharedFiles = $state<any[]>([]);
  let isLoading = $state(true);
  let isConnecting = $state(false);
  let sentMessages = $state<Set<string>>(new Set()); // Track sent messages to avoid duplicates

  // Filter rooms based on search query
  $effect(() => {
    if (searchQuery) {
      filteredRooms = chatRooms.filter((room) =>
        (room.title || "").toLowerCase().includes(searchQuery.toLowerCase())
      );
    } else {
      filteredRooms = chatRooms;
    }
  });

  let filteredRooms = $state<any[]>([]);

  // Load chat rooms when component mounts
  onMount(async () => {
    if (browser) {
      try {
        const response = await fetch("/api/chat");
        if (response.ok) {
          const data = await response.json();
          chatRooms = data || [];
          filteredRooms = [...chatRooms];
          console.log("Loaded chat rooms:", chatRooms);
        } else {
          console.error("Failed to load chat rooms");
        }
      } catch (error) {
        console.error("Error loading chat rooms:", error);
      } finally {
        isLoading = false;
      }
    }
  });

  // Hàm tìm tên người dùng dựa trên ID
  function findUserNameById(userId: string) {
    if (!teamMembers?.length) return null;
    const member = teamMembers.find((m) => m.id === userId);
    return member ? member.name : null;
  }

  // Connect to WebSocket
  function connectWebSocket() {
    if (!selectedRoom || !browser || isConnecting) return;

    isConnecting = true;

    // Đóng kết nối hiện tại nếu có
    if (socket) {
      try {
        socket.close(1000, "Switching rooms");
      } catch (error) {
        console.error("Error closing existing WebSocket:", error);
      }
      socket = null;
    }

    const projectID = selectedRoom.project_id;
    const userID = data.id;

    console.log("WebSocket params:", { projectID, userID });

    const protocol = window.location.protocol === "https:" ? "wss" : "ws";
    const host = window.location.host.includes("skillforge.ikniz.site")
      ? window.location.host
      : "localhost:8080";

    const wsUrl = `${protocol}://${host}/ws/chats/${projectID}/${userID}`;
    console.log("Connecting to WebSocket:", wsUrl);

    try {
      socket = new WebSocket(wsUrl);

      socket.onopen = () => {
        console.log("WebSocket connection established for chat");
        isConnecting = false; // Đánh dấu kết nối hoàn tất
      };

      socket.onmessage = (event) => {
        try {
          const newMessage = JSON.parse(event.data);
          console.log("Received message:", newMessage);

          // Xử lý đặc biệt cho trường hợp tin nhắn của chính người dùng
          const isCurrentUserMessage = newMessage.sender_id === data.id;

          // Map các trường để phù hợp với cấu trúc Message từ server
          const processedMessage = {
            id: newMessage.id,
            content: newMessage.content,
            timestamp: newMessage.created_at,
            sender_id: newMessage.sender_id,
            user_id: newMessage.sender_id, // Đồng bộ với UI
            group_id: newMessage.group_id,
            user_name: isCurrentUserMessage
              ? "You"
              : findUserNameById(newMessage.sender_id) || "Unknown User",
          };

          // Thêm tin nhắn mới vào UI, tránh trùng lặp
          const isDuplicate = messages.some(
            (msg) =>
              msg.content === processedMessage.content &&
              msg.sender_id === processedMessage.sender_id &&
              // Nếu đã có tin nhắn tương tự trong vòng 5 giây, coi là trùng lặp
              (msg.id === processedMessage.id ||
                (msg.timestamp &&
                  Math.abs(
                    new Date(msg.timestamp).getTime() -
                      new Date(processedMessage.timestamp).getTime()
                  ) < 5000))
          );

          if (!isDuplicate) {
            messages = [...messages, processedMessage];

            // Auto-scroll to bottom when new message arrives
            setTimeout(() => {
              const chatContainer = document.getElementById("chat-messages");
              if (chatContainer) {
                chatContainer.scrollTop = chatContainer.scrollHeight;
              }
            }, 50);
          }
        } catch (error) {
          console.error("Error parsing message:", error);
        }
      };

      socket.onerror = (error) => {
        console.error("WebSocket error:", error);
        isConnecting = false; // Reset trạng thái khi gặp lỗi
      };

      socket.onclose = (event) => {
        console.log(
          `WebSocket connection closed: ${event.code} ${event.reason}`
        );
        isConnecting = false; // Reset trạng thái khi kết nối đóng
      };
    } catch (error) {
      console.error("Failed to create WebSocket connection:", error);
      isConnecting = false; // Reset trạng thái khi gặp lỗi
    }
  }

  // Load messages for a selected room
  async function selectRoom(room: any) {
    // Nếu đang chọn phòng đang kết nối, không làm gì cả
    if (selectedRoom?.project_id === room.project_id) return;

    // Đóng kết nối trước khi thay đổi phòng
    if (socket) {
      try {
        socket.close(1000, "Switching rooms");
      } catch (error) {
        console.error("Error closing WebSocket:", error);
      }
      socket = null;
    }

    // Cập nhật phòng được chọn
    selectedRoom = room;
    console.log("Selected room:", room);

    // Reset tin nhắn đã gửi khi chuyển phòng
    sentMessages.clear();

    try {
      const response = await fetch(`/api/chat/${room.project_id}`);
      if (response.ok) {
        const data = await response.json();
        console.log("Room data:", data);

        // Map lại các tin nhắn để đảm bảo đúng định dạng
        const processedMessages = (data.messages || []).map((msg: any) => {
          const isCurrentUserMessage = msg.sender_id === data.id;
          return {
            content: msg.content,
            timestamp: msg.created_at || msg.timestamp,
            sender_id: msg.sender_id,
            user_id: msg.sender_id,
            user_name: isCurrentUserMessage
              ? "You"
              : findUserNameById(msg.sender_id) || "Unknown User",
          };
        });

        messages = processedMessages;
        teamMembers = data.members || [];
        console.log("Room data loaded, team members:", teamMembers);

        // Sau khi tải dữ liệu xong, mới kết nối WebSocket
        connectWebSocket();

        // Auto-scroll to bottom when loading messages
        setTimeout(() => {
          const chatContainer = document.getElementById("chat-messages");
          if (chatContainer) {
            chatContainer.scrollTop = chatContainer.scrollHeight;
          }
        }, 50);
      } else {
        console.error("Failed to load chat messages");
      }
    } catch (error) {
      console.error("Error loading chat messages:", error);
    }
  }

  // Send a message
  function sendMessage() {
    if (!messageInput.trim() || !socket) {
      return;
    }

    // Chỉ gửi nếu socket đã sẵn sàng
    if (socket.readyState === WebSocket.OPEN) {
      const messageContent = messageInput.trim();

      // Cấu trúc message để gửi đi - chỉ gửi content vì server không sử dụng user_name
      const messageToSend = {
        content: messageContent,
      };

      try {
        socket.send(JSON.stringify(messageToSend));
        console.log("Message sent:", messageToSend);

        // Ghi lại ID tin nhắn đã gửi để tránh hiển thị lặp khi nhận lại từ server
        const tempId = `${messageContent}-${data.id}-${new Date().toISOString()}`;
        sentMessages.add(tempId);

        // Không tự thêm tin nhắn vào UI nữa, đợi nhận từ server
        messageInput = ""; // Clear input after sending
      } catch (error) {
        console.error("Error sending message:", error);
      }
    } else {
      console.warn("WebSocket not ready, message not sent");
    }
  }

  // Handle Enter key to send message
  function handleKeyDown(event: KeyboardEvent) {
    if (event.key === "Enter" && !event.shiftKey) {
      event.preventDefault();
      sendMessage();
    }
  }

  // Format message time
  function formatTime(timestamp: string) {
    if (!timestamp) return "";

    try {
      const date = new Date(timestamp);
      return date.toLocaleTimeString([], {
        hour: "2-digit",
        minute: "2-digit",
      });
    } catch (e) {
      console.error("Error formatting time:", e);
      return "";
    }
  }

  // Format date for message groups
  function formatDate(timestamp: string) {
    if (!timestamp) return "";

    try {
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
          year: "numeric",
        });
      }
    } catch (e) {
      console.error("Error formatting date:", e);
      return "";
    }
  }

  // Clean up WebSocket connection on component destroy
  onDestroy(() => {
    if (socket) {
      try {
        // Đóng với code 1000 (normal closure)
        socket.close(1000, "Component unmounted");
      } catch (e) {
        console.error("Error closing WebSocket on destroy:", e);
      }
      socket = null;
    }
  });
</script>

<main class="flex-1 flex flex-col ml-64 pr-4 pl-4 h-[calc(100vh-20px)] pt-4">
  <div class="flex space-x-4 flex-1">
    <!-- Chat Rooms List -->
    <div class="card p-3 w-64 flex flex-col">
      <div class="flex items-center justify-between mb-3">
        <h3 class="text-base font-semibold">Rooms</h3>
        <button class="text-[#6b48ff] hover:text-[#5a3de6]"></button>
      </div>

      <!-- Search Box -->
      <div class="relative mb-3">
        <input
          type="text"
          placeholder="Search rooms..."
          bind:value={searchQuery}
          class="w-full p-2 pl-8 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-[#6b48ff]"
        />
        <svg
          class="w-4 h-4 absolute left-2 top-2.5 text-gray-400"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          />
        </svg>
      </div>

      <!-- Room List -->
      <div class="flex-1 overflow-y-auto">
        {#if isLoading}
          <div class="flex justify-center items-center h-32">
            <div
              class="animate-spin rounded-full h-8 w-8 border-b-2 border-[#6b48ff]"
            ></div>
          </div>
        {:else if filteredRooms.length === 0}
          <div class="text-center p-4 text-gray-500">No chat rooms found</div>
        {:else}
          <div class="space-y-1">
            {#each filteredRooms as room}
              <div
                class={`chat-room p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center ${selectedRoom?.project_id === room.project_id ? "active" : ""}`}
                onclick={() => selectRoom(room)}
              >
                <span class="text-sm font-medium">{room.title}</span>
              </div>
            {/each}
          </div>
        {/if}
      </div>
    </div>

    <!-- Chat Window -->
    <div class="card p-4 flex-1 flex flex-col">
      {#if !selectedRoom}
        <!-- No room selected state -->
        <div
          class="flex-1 flex items-center justify-center bg-gray-100 p-3 rounded mb-4"
        >
          <div class="text-center p-6">
            <svg
              class="w-16 h-16 mx-auto text-gray-400 mb-4"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"
              />
            </svg>
            <h3 class="text-lg font-medium text-gray-700 mb-2">
              No conversation selected
            </h3>
            <p class="text-gray-500 mb-4">
              Choose a room from the list to start chatting
            </p>
          </div>
        </div>
      {:else}
        <!-- Chat Header -->
        <div
          class="flex justify-between items-center p-2 border-b border-gray-200 mb-3"
        >
          <div>
            <h3 class="font-medium">
              {selectedRoom.title || selectedRoom.name || "Chat Room"}
            </h3>
            <p class="text-xs text-gray-500">
              {teamMembers?.length || 0} members
            </p>
          </div>
        </div>
        <!-- Chat Messages Area -->
        <div id="chat-messages" class="flex-1 overflow-y-auto mb-4 space-y-4">
          {#if messages.length === 0}
            <div class="text-center p-6 text-gray-500">
              No messages yet. Start the conversation!
            </div>
          {:else}
            <!-- Group messages by date -->
            {#each messages as message, i}
              <!-- Date separator -->
              {#if i === 0 || formatDate(messages[i - 1].timestamp) !== formatDate(message.timestamp)}
                <div class="flex justify-center">
                  <div
                    class="bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded"
                  >
                    {formatDate(message.timestamp)}
                  </div>
                </div>
              {/if}

              <!-- Message item -->
              <div
                class={`flex ${message.user_id === data.id || message.sender_id === data.id ? "justify-end" : "justify-start"}`}
              >
                <div
                  class={`max-w-[70%] ${message.user_id === data.id || message.sender_id === data.id ? "bg-[#6b48ff] text-white" : "bg-gray-100 text-gray-800"} rounded-lg px-3 py-2`}
                >
                  {#if message.user_id !== data.id && message.sender_id !== data.id}
                    <div class="text-xs font-semibold mb-1">
                      {message.user_name || "Unknown User"}
                    </div>
                  {:else}
                    <div class="text-xs font-semibold mb-1 text-purple-100">
                      You
                    </div>
                  {/if}
                  <p>{message.content}</p>
                  <div
                    class={`text-right text-xs ${message.user_id === data.id || message.sender_id === data.id ? "text-purple-200" : "text-gray-500"} mt-1`}
                  >
                    {formatTime(message.timestamp || message.created_at)}
                  </div>
                </div>
              </div>
            {/each}
          {/if}
        </div>
      {/if}

      <!-- Message Input Area -->
      <div class="flex space-x-2">
        <input
          type="text"
          bind:value={messageInput}
          onkeydown={handleKeyDown}
          class="flex-1 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff] disabled:bg-gray-100 disabled:cursor-not-allowed"
          placeholder={selectedRoom
            ? "Type your message..."
            : "Select a room to start messaging..."}
          disabled={!selectedRoom}
        />
        <button
          class={`btn ${!selectedRoom ? "opacity-50 cursor-not-allowed" : ""}`}
          disabled={!selectedRoom}
          onclick={sendMessage}
        >
          Send
        </button>
        <label
          class={`btn-secondary flex items-center ${!selectedRoom ? "opacity-50 cursor-not-allowed" : "cursor-pointer"}`}
        >
          <svg
            class="w-5 h-5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"
            />
          </svg>
          <input
            type="file"
            class="hidden"
            accept="image/*"
            disabled={!selectedRoom}
          />
        </label>
      </div>
      <button
        class={`btn-secondary mt-2 w-full flex items-center justify-center ${!selectedRoom ? "opacity-50 cursor-not-allowed" : ""}`}
        disabled={!selectedRoom}
      >
        <svg
          class="w-5 h-5 mr-1"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"
          />
        </svg>
        Start Video Call
      </button>
    </div>

    <!-- Team Members -->
    <!-- Team Members -->
    <div class="card p-4 w-64 flex flex-col">
      <h3 class="text-base font-semibold mb-3">
        Team Members ({teamMembers.length})
      </h3>
      <div class="flex-1 overflow-y-auto space-y-3">
        {#if !selectedRoom}
          <div class="text-center p-4 text-gray-500">
            Select a room to see members
          </div>
        {:else if !teamMembers || teamMembers.length === 0}
          <div class="text-center p-4 text-gray-500">No team members found</div>
        {:else}
          {#each teamMembers as member (member.id)}
            <div class="flex items-center p-2 hover:bg-gray-50 rounded">
              <!-- Avatar with fallback -->
              <div class="relative w-10 h-10 mr-3 flex-shrink-0">
                {#if member.avatar_name}
                  <img
                    class="w-10 h-10 rounded-full object-cover border border-gray-200"
                    src={`/api/avatars/${member.id}`}
                    alt={member.name}
                  />
                {:else}
                  <div
                    class="w-10 h-10 rounded-full bg-[#6b48ff] flex items-center justify-center text-white font-medium"
                  >
                    {member.name.charAt(0).toUpperCase()}
                  </div>
                {/if}

                <!-- Online status indicator -->
                <span
                  class="absolute bottom-0 right-0 w-3 h-3 rounded-full bg-green-400 border-2 border-white"
                ></span>
              </div>

              <!-- User info -->
              <div class="overflow-hidden">
                <p class="text-sm font-medium truncate">{member.name}</p>
                <div class="flex items-center">
                  {#if member.title}
                    <p class="text-xs text-gray-500 truncate">{member.title}</p>
                  {:else}
                    <p class="text-xs text-gray-500 capitalize">
                      {member.role || "Member"}
                    </p>
                  {/if}
                </div>

                <!-- User actions - only show for current user -->
                {#if member.id === data.id}
                  <span
                    class="inline-block text-xs bg-blue-100 text-blue-800 px-1.5 py-0.5 rounded-full mt-1"
                    >You</span
                  >
                {/if}
              </div>
            </div>
          {/each}
        {/if}
      </div>

      <!-- Phân chia rõ ràng -->
      <hr class="my-3 border-gray-200" />

      <!-- Shared Files section -->
      <h3 class="text-base font-semibold mb-3">Shared Files</h3>
      <div class="space-y-2 flex-1 overflow-y-auto">
        {#if !selectedRoom}
          <div class="text-center p-4 text-gray-500">
            Select a room to see files
          </div>
        {:else if sharedFiles.length === 0}
          <div class="text-center p-4 text-gray-500">No shared files yet</div>
          <button
            class="btn-secondary w-full text-sm mt-2 flex items-center justify-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4 mr-1"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 4v16m8-8H4"
              />
            </svg>
            Share a file
          </button>
        {:else}
          {#each sharedFiles as file}
            <div class="p-2 bg-gray-50 rounded text-sm flex items-center">
              <svg
                class="w-4 h-4 text-blue-500 mr-2 flex-shrink-0"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                />
              </svg>
              <div class="overflow-hidden">
                <p class="truncate">{file.name}</p>
                <p class="text-xs text-gray-400">
                  Shared by {file.shared_by_name || "Unknown"}
                </p>
              </div>
            </div>
          {/each}
        {/if}
      </div>
    </div>
  </div>
</main>

<style>
  .chat-room.active {
    background-color: rgba(107, 72, 255, 0.1);
    border-left: 3px solid #6b48ff;
  }

  .chat-room.active:hover {
    background-color: rgba(107, 72, 255, 0.15);
  }

  .btn {
    background-color: #6b48ff;
    color: white;
    padding: 0.5rem 1rem;
    border-radius: 0.375rem;
    font-weight: 500;
    transition-property: background-color, border-color, color, fill, stroke;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }

  .btn:hover {
    background-color: #5a3cd9;
  }

  .btn-secondary {
    background-color: white;
    color: #4b5563;
    padding: 0.5rem 1rem;
    border: 1px solid #d1d5db;
    border-radius: 0.375rem;
    font-weight: 500;
    transition-property: background-color, border-color, color, fill, stroke;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }

  .btn-secondary:hover {
    background-color: #f3f4f6;
  }

  .card {
    background-color: white;
    border-radius: 0.5rem;
    box-shadow:
      0 1px 3px 0 rgba(0, 0, 0, 0.1),
      0 1px 2px 0 rgba(0, 0, 0, 0.06);
  }

  #chat-messages {
    height: calc(100% - 60px);
    scrollbar-width: thin;
    scrollbar-color: #d1d5db transparent;
  }

  #chat-messages::-webkit-scrollbar {
    width: 6px;
  }

  #chat-messages::-webkit-scrollbar-track {
    background: transparent;
  }

  #chat-messages::-webkit-scrollbar-thumb {
    background-color: #d1d5db;
    border-radius: 3px;
  }
</style>
