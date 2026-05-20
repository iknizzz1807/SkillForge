<script lang="ts">
  import { onMount } from "svelte";

  let chatRooms: any[] = [];
  let messages: any[] = [];
  let teamMembers: any[] = [];
  let sharedFiles: any[] = [];

  let selectedRoom: any = null;
  let messageInput = "";
  let searchQuery = "";
  let filteredRooms: any[] = [];

  import { PUBLIC_WS_URL } from '$env/static/public';
  import { get } from 'svelte/store';
  import { user } from '$lib/stores/auth'; // assuming standard auth store or just get token from localStorage

  onMount(async () => {
    try {
      const token = localStorage.getItem('token') || "";
      // The API path is on the backend, usually prefixed with PUBLIC_API_URL or it's directly proxy.
      // Let's assume there is a fetch wrapper or we just call the full URL or relative.
      const API_URL = "http://localhost:8080"; // fallback if no env
      const res = await fetch(`${API_URL}/api/chats`, {
        headers: { 'Authorization': `Bearer ${token}` }
      });
      if (res.ok) {
        chatRooms = await res.json() || [];
        filteredRooms = chatRooms;
        if (chatRooms.length > 0) {
          selectRoom(chatRooms[0]);
        }
      }
    } catch (e) {
      console.error("Failed to fetch chat rooms", e);
    }
  });

  async function loadRoomDetails(roomId: string) {
    try {
      const token = localStorage.getItem('token') || "";
      const API_URL = "http://localhost:8080";
      const res = await fetch(`${API_URL}/api/chats/${roomId}`, {
        headers: { 'Authorization': `Bearer ${token}` }
      });
      if (res.ok) {
        const data = await res.json();
        messages = data.messages || [];
        teamMembers = data.members || [];
      }
    } catch(e) {
      console.error("Failed to load room details", e);
    }
  }

  // Filter rooms based on search
  $effect(() => {
    if (searchQuery) {
      filteredRooms = chatRooms.filter((room) =>
        room.title.toLowerCase().includes(searchQuery.toLowerCase())
      );
    } else {
      filteredRooms = chatRooms;
    }
  });

  // Select a room
  function selectRoom(room: any) {
    selectedRoom = room;
    // Mark room as read by resetting unread count
    if (room.unread !== undefined) {
        room.unread = 0;
    }

    loadRoomDetails(room.id);

    // Auto-scroll to bottom when changing rooms
    setTimeout(() => {
      const chatContainer = document.getElementById("chat-messages");
      if (chatContainer) {
        chatContainer.scrollTop = chatContainer.scrollHeight;
      }
    }, 50);
  }

  // Send message
  function sendMessage() {
    if (!messageInput.trim()) return;

    // Add message to the chat
    messages = [...messages, {
      id: `m${messages.length + 1}`,
      content: messageInput.trim(),
      sender_id: "current_user",
      user_name: "You",
      avatar: "/images/avatar.jpg",
      timestamp: new Date().toISOString(),
      isFile: false,
    }];

    // Clear input
    messageInput = "";

    // Auto-scroll to the newest message
    setTimeout(() => {
      const chatContainer = document.getElementById("chat-messages");
      if (chatContainer) {
        chatContainer.scrollTop = chatContainer.scrollHeight;
      }
    }, 50);
  }

  // Format message time
  function formatTime(timestamp: string | Date) {
    const date = new Date(timestamp);
    return date.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
  }

  // Format date for message groups
  function formatDate(timestamp: string | Date) {
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

  // Handle Enter key to send message
  function handleKeyDown(event: KeyboardEvent) {
    if (event.key === "Enter" && !event.shiftKey) {
      event.preventDefault();
      sendMessage();
    }
  }

  // Scroll to bottom on mount
  onMount(() => {
    const chatContainer = document.getElementById("chat-messages");
    if (chatContainer) {
      chatContainer.scrollTop = chatContainer.scrollHeight;
    }
  });
</script>

<main class="flex-1 flex flex-col ml-64 p-4 h-screen">
  <div class="flex gap-4 flex-1 h-full overflow-hidden">
    <!-- Chat Rooms List - Thu gọn -->
    <div class="card p-3 w-64 flex flex-col flex-shrink-0">
      <div class="flex items-center justify-between mb-2">
        <h3 class="text-sm font-semibold text-gray-800">Conversations</h3>
        <div class="flex space-x-1">
          <button class="p-1 rounded-full hover:bg-gray-100 text-gray-500">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                d="M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 5a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V5zM14 11a1 1 0 011 1v1h1a1 1 0 110 2h-1v1a1 1 0 11-2 0v-1h-1a1 1 0 110-2h1v-1a1 1 0 011-1z"
              />
            </svg>
          </button>
        </div>
      </div>

      <!-- Search Box - Thu gọn -->
      <div class="relative mb-2">
        <input
          type="text"
          placeholder="Search conversations..."
          bind:value={searchQuery}
          class="w-full p-1.5 pl-7 text-xs border border-gray-200 rounded-lg focus:outline-none focus:ring-1 focus:ring-[#6b48ff] bg-gray-50"
        />
        <svg
          class="w-3.5 h-3.5 absolute left-2 top-2 text-gray-400"
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

      <!-- Room List - Thu gọn -->
      <div class="flex-1 overflow-y-auto custom-scrollbar">
        <div class="space-y-1">
          {#each filteredRooms as room}
            <div
              class={`p-2 rounded-lg cursor-pointer hover:bg-gray-50 flex items-center transition-colors ${
                selectedRoom.id === room.id ? "bg-[#6b48ff] bg-opacity-10 border-l-2 border-[#6b48ff]" : "border-l-2 border-transparent"
              }`}
              on:click={() => selectRoom(room)}
            >
              <!-- Room icon -->
              <div
                class="w-8 h-8 rounded-full bg-gradient-to-br from-purple-500 to-indigo-600 flex items-center justify-center flex-shrink-0 mr-2"
              >
                <span class="text-white font-medium text-xs"
                  >{room.title.substring(0, 2).toUpperCase()}</span
                >
              </div>

              <!-- Room info -->
              <div class="flex-1 overflow-hidden">
                <div class="flex justify-between">
                  <span class="text-xs font-medium text-gray-800 truncate"
                    >{room.title}</span
                  >
                  <span class="text-[10px] text-gray-500"
                    >{room.lastMessageTime}</span
                  >
                </div>
                <div class="flex justify-between items-center">
                  <p class="text-[10px] text-gray-500 truncate">
                    {room.lastMessage}
                  </p>
                  {#if room.unread > 0}
                    <span
                      class="flex-shrink-0 bg-[#6b48ff] text-white text-[10px] rounded-full h-4 w-4 flex items-center justify-center ml-1"
                      >{room.unread}</span
                    >
                  {/if}
                </div>
              </div>
            </div>
          {/each}
        </div>
      </div>
    </div>

    <!-- Chat Window - Thu gọn -->
    <div class="card p-0 flex-1 flex flex-col overflow-hidden">
      <!-- Chat Header - Thu gọn -->
      <div
        class="flex justify-between items-center px-3 py-2 border-b border-gray-100"
      >
        <div class="flex items-center">
          <div
            class="w-8 h-8 rounded-full bg-gradient-to-br from-purple-500 to-indigo-600 flex items-center justify-center mr-2"
          >
            <span class="text-white font-medium text-xs"
              >{selectedRoom.title.substring(0, 2).toUpperCase()}</span
            >
          </div>
          <div>
            <h3 class="text-sm font-medium text-gray-800">
              {selectedRoom.title}
            </h3>
            <div class="flex items-center">
              <span class="w-1.5 h-1.5 rounded-full bg-green-500 mr-1"></span>
              <p class="text-[10px] text-gray-500">
                {teamMembers.filter((m) => m.status === "online").length} online
              </p>
            </div>
          </div>
        </div>

        <div class="flex items-center space-x-1">
          <button class="p-1 rounded-full hover:bg-gray-100 text-gray-500">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>
          </button>
          <button class="p-1 rounded-full hover:bg-gray-100 text-gray-500">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"
              />
            </svg>
          </button>
        </div>
      </div>

      <!-- Chat Messages Area -->
      <div
        id="chat-messages"
        class="flex-1 overflow-y-auto p-2 space-y-2 bg-gradient-to-b from-gray-50 to-white custom-scrollbar"
      >
        <!-- Group messages by date -->
        {#each messages as message, i}
          <!-- Date separator -->
          {#if i === 0 || formatDate(messages[i - 1].timestamp) !== formatDate(message.timestamp)}
            <div class="flex justify-center my-1">
              <div
                class="bg-white text-gray-500 text-[10px] px-2 py-0.5 rounded-full shadow-sm border border-gray-100"
              >
                {formatDate(message.timestamp)}
              </div>
            </div>
          {/if}

          <!-- Message item -->
          <div
            class={`flex ${message.sender_id === "current_user" ? "justify-end" : "justify-start"}`}
          >
            {#if message.sender_id !== "current_user"}
              <img
                src={message.avatar}
                alt={message.user_name}
                class="w-6 h-6 rounded-full mr-1 mt-1 border border-white shadow-sm"
              />
            {/if}

            <div
              class={`max-w-[65%] ${
                message.sender_id === "current_user"
                  ? "bg-[#6b48ff] text-white rounded-tl-lg rounded-tr-lg rounded-bl-lg"
                  : "bg-white border border-gray-100 shadow-sm text-gray-800 rounded-tl-lg rounded-tr-lg rounded-br-lg"
              } px-2.5 py-1.5`}
            >
              {#if message.sender_id !== "current_user"}
                <div class="text-[10px] font-semibold text-gray-700">
                  {message.user_name}
                </div>
              {/if}

              {#if message.isFile}
                <!-- File message -->
                <div
                  class="flex items-center bg-white bg-opacity-10 rounded p-1 border border-white border-opacity-20"
                >
                  <div class="bg-white bg-opacity-20 p-1 rounded mr-2">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3.5 w-3.5"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                      />
                    </svg>
                  </div>
                  <div>
                    <p
                      class={message.sender_id === "current_user"
                        ? "text-white text-xs"
                        : "text-gray-800 text-xs"}
                    >
                      {message.fileName}
                    </p>
                    <p
                      class={message.sender_id === "current_user"
                        ? "text-purple-200 text-[10px]"
                        : "text-gray-500 text-[10px]"}
                    >
                      {message.fileSize}
                    </p>
                  </div>
                  <button
                    class={`ml-2 p-1 rounded-full ${
                      message.sender_id === "current_user"
                        ? "bg-white bg-opacity-20 hover:bg-opacity-30"
                        : "bg-gray-100 hover:bg-gray-200"
                    }`}
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3 w-3"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"
                      />
                    </svg>
                  </button>
                </div>
              {:else}
                <!-- Text message -->
                <p class="text-xs">{message.content}</p>
              {/if}

              <div
                class={`text-right text-[10px] ${
                  message.sender_id === "current_user"
                    ? "text-purple-200"
                    : "text-gray-400"
                } mt-0.5`}
              >
                {formatTime(message.timestamp)}
              </div>
            </div>

            {#if message.sender_id === "current_user"}
              <img
                src={message.avatar || "/images/avatar-placeholder.png"}
                alt={message.user_name}
                class="w-6 h-6 rounded-full ml-1 mt-1 border border-white shadow-sm"
              />
            {/if}
          </div>
        {/each}
      </div>

      <!-- Message Input Area -->
      <div class="p-2 border-t border-gray-100">
        <div
          class="flex space-x-1 items-center bg-gray-50 p-1 rounded-lg border border-gray-200"
        >
          <button class="p-1 rounded-full hover:bg-gray-200 text-gray-500">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zM7 9a1 1 0 100-2 1 1 0 000 2zm7-1a1 1 0 11-2 0 1 1 0 012 0zm-7.536 5.879a1 1 0 001.415 0 3 3 0 014.242 0 1 1 0 001.415-1.415 5 5 0 00-7.072 0 1 1 0 000 1.415z"
                clip-rule="evenodd"
              />
            </svg>
          </button>

          <button class="p-1 rounded-full hover:bg-gray-200 text-gray-500">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-8 3 6 2-4 3 6z"
                clip-rule="evenodd"
              />
            </svg>
          </button>

          <input
            type="text"
            bind:value={messageInput}
            on:keydown={handleKeyDown}
            class="flex-1 p-1.5 text-xs bg-transparent border-none focus:outline-none"
            placeholder="Type your message..."
          />

          <button
            class="bg-[#6b48ff] text-white p-1.5 rounded-full transition-colors hover:bg-[#5a3de6] disabled:bg-gray-300 disabled:cursor-not-allowed"
            on:click={sendMessage}
            disabled={!messageInput.trim()}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Team Members & Files Panel - Thu gọn -->
    <div class="card p-3 w-64 flex flex-col flex-shrink-0">
      <!-- Team Members Section -->
      <div class="mb-2">
        <h3
          class="text-xs font-semibold text-gray-800 mb-1.5 flex items-center"
        >
          Team Members
          <span
            class="ml-1 text-[10px] bg-[#6b48ff] text-white px-1 rounded-full"
            >{teamMembers.length}</span
          >
        </h3>

        <div
          class="space-y-1 mb-1.5 overflow-y-auto custom-scrollbar max-h-[40vh]"
        >
          {#each teamMembers as member}
            <div
              class="flex items-center p-1 hover:bg-gray-50 rounded-lg transition-colors"
            >
              <div class="relative">
                <img
                  src={member.avatar || "/images/avatar-placeholder.png"}
                  alt={member.name}
                  class="w-6 h-6 rounded-full object-cover border border-white shadow-sm"
                />
                <div
                  class={`absolute -bottom-0.5 -right-0.5 w-2 h-2 rounded-full ${
                    member.status === "online"
                      ? "bg-green-500"
                      : member.status === "away"
                        ? "bg-yellow-500"
                        : "bg-gray-300"
                  } border border-white`}
                ></div>
              </div>

              <div class="ml-1.5 flex-1 overflow-hidden">
                <div class="flex justify-between items-center">
                  <p class="text-xs font-medium truncate">{member.name}</p>
                  {#if member.isCurrentUser}
                    <span
                      class="text-[8px] bg-blue-100 text-blue-700 px-1 rounded-full"
                      >You</span
                    >
                  {/if}
                </div>
                <p class="text-[9px] text-gray-500 truncate">{member.title}</p>
              </div>
            </div>
          {/each}
        </div>
      </div>

      <div class="border-t border-gray-100 pt-2">
        <!-- Shared Files Section -->
        <h3
          class="text-xs font-semibold text-gray-800 mb-1.5 flex items-center justify-between"
        >
          <span>Shared Files</span>
          <button class="text-[9px] text-[#6b48ff] hover:text-[#5a3de6]"
            >View all</button
          >
        </h3>
        <div class="space-y-1 overflow-y-auto custom-scrollbar max-h-[41vh]">
          {#each sharedFiles as file}
            <div
              class="p-1.5 bg-gray-50 rounded-lg border border-gray-100 hover:border-gray-200 transition-colors"
            >
              <div class="flex items-center mb-0.5">
                <div
                  class={`p-1 rounded mr-1.5 ${
                    file.type === "document"
                      ? "bg-blue-100 text-blue-600"
                      : file.type === "design"
                        ? "bg-purple-100 text-purple-600"
                        : file.type === "spreadsheet"
                          ? "bg-green-100 text-green-600"
                          : "bg-gray-100 text-gray-600"
                  }`}
                >
                  {#if file.type === "document"}
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3 w-3"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z"
                        clip-rule="evenodd"
                      />
                    </svg>
                  {:else if file.type === "design"}
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3 w-3"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M4 2a2 2 0 00-2 2v11a3 3 0 106 0V4a2 2 0 00-2-2H4zm1 14a1 1 0 100-2 1 1 0 000 2zm5-1.757l4.9-4.9a2 2 0 000-2.828L13.485 5.1a2 2 0 00-2.828 0L10 5.757v8.486zM16 18H9.071l6-6H16a2 2 0 012 2v2a2 2 0 01-2 2z"
                        clip-rule="evenodd"
                      />
                    </svg>
                  {:else if file.type === "spreadsheet"}
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3 w-3"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M5 4a3 3 0 00-3 3v6a3 3 0 003 3h10a3 3 0 003-3V7a3 3 0 00-3-3H5zm-1 9v-1h5v2H5a1 1 0 01-1-1zm7 1h4a1 1 0 001-1v-1h-5v2zm0-4h5V8h-5v2zM9 8H4v2h5V8z"
                        clip-rule="evenodd"
                      />
                    </svg>
                  {:else}
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3 w-3"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M8 4a3 3 0 00-3 3v4a5 5 0 0010 0V7a1 1 0 112 0v4a7 7 0 11-14 0V7a5 5 0 0110 0v4a3 3 0 11-6 0V7a1 1 0 012 0v4a1 1 0 102 0V7a3 3 0 00-3-3z"
                        clip-rule="evenodd"
                      />
                    </svg>
                  {/if}
                </div>
                <div class="flex-1 overflow-hidden">
                  <p class="text-[10px] font-medium truncate">{file.name}</p>
                </div>
                <button
                  class="p-0.5 text-gray-400 hover:text-gray-600 rounded-full hover:bg-gray-200 ml-1"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-3 w-3"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </button>
              </div>
              <div class="flex justify-between text-[9px] text-gray-500">
                <span>{file.size}</span>
                <span>{file.date}</span>
              </div>
            </div>
          {/each}
        </div>

        <div class="mt-2">
          <button
            class="btn-secondary w-full flex items-center justify-center text-xs py-1.5"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-3 w-3 mr-1"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM6.293 6.707a1 1 0 010-1.414l3-3a1 1 0 011.414 0l3 3a1 1 0 01-1.414 1.414L11 5.414V13a1 1 0 11-2 0V5.414L7.707 6.707a1 1 0 01-1.414 0z"
                clip-rule="evenodd"
              />
            </svg>
            Upload File
          </button>
        </div>
      </div>
    </div>
  </div>
</main>
