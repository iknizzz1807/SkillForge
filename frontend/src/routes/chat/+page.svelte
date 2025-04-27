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

  // Connect to WebSocket when a room is selected
  $effect(() => {
    if (selectedRoom && browser) {
      connectWebSocket();
    }
  });

  // Filter rooms based on search query
  $effect(() => {
    if (searchQuery) {
      filteredRooms = chatRooms.filter((room) =>
        room.name.toLowerCase().includes(searchQuery.toLowerCase())
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
          // [
          //     {
          //         "project_id": "000d764b-4819-479a-b6e8-b620817dc6d3",
          //         "title": "Make a website",
          //         "created_at": "0001-01-01T00:00:00Z"
          //     },
          //     {
          //         "project_id": "000d764b-4819-479a-b6e8-b620817dc6d3",
          //         "title": "Make a website",
          //         "created_at": "0001-01-01T00:00:00Z"
          //     }
          // ]
          chatRooms = data || [];
          filteredRooms = [...chatRooms];
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

  // Connect to WebSocket
  function connectWebSocket() {
    if (!selectedRoom || !browser) return;

    // Close existing connection
    if (socket) {
      socket.close();
    }

    const projectID = selectedRoom.project_id;
    const userID = data.id; // Assuming user ID is in page data

    const protocol = window.location.protocol === "https:" ? "wss" : "ws";
    const host = window.location.host.includes("skillforge.ikniz.site")
      ? window.location.host
      : "localhost:8080";

    const wsUrl = `${protocol}://${host}/ws/chats/${projectID}/${userID}`;

    socket = new WebSocket(wsUrl);

    socket.onopen = () => {
      console.log("WebSocket connection established for chat");
    };

    socket.onmessage = (event) => {
      try {
        const newMessage = JSON.parse(event.data);
        messages = [...messages, newMessage];

        // Auto-scroll to bottom when new message arrives
        setTimeout(() => {
          const chatContainer = document.getElementById("chat-messages");
          if (chatContainer) {
            chatContainer.scrollTop = chatContainer.scrollHeight;
          }
        }, 50);
      } catch (error) {
        console.error("Error parsing message:", error);
      }
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    socket.onclose = () => {
      console.log("WebSocket connection closed");
    };
  }

  // Load messages for a selected room
  async function selectRoom(room: any) {
    selectedRoom = room;

    try {
      const response = await fetch(`/api/chat/${room.project_id}`);
      if (response.ok) {
        const data = await response.json();
        messages = data.messages;
        teamMembers = data.members;

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
    if (
      !messageInput.trim() ||
      !socket ||
      socket.readyState !== WebSocket.OPEN
    ) {
      return;
    }

    const message = {
      content: messageInput,
    };

    socket.send(JSON.stringify(message));
    messageInput = ""; // Clear input after sending
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

    const date = new Date(timestamp);
    return date.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
  }

  // Format date for message groups
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
        year: "numeric",
      });
    }
  }

  // Clean up WebSocket connection on component destroy
  onDestroy(() => {
    if (socket) {
      socket.close();
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
                class={`chat-room p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center ${selectedRoom?.id === room.id ? "active" : ""}`}
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
            <h3 class="font-medium">{selectedRoom.name}</h3>
            <p class="text-xs text-gray-500">{teamMembers.length} members</p>
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
                class={`flex ${message.user_id === data.id ? "justify-end" : "justify-start"}`}
              >
                <div
                  class={`max-w-[70%] ${message.user_id === data.id ? "bg-[#6b48ff] text-white" : "bg-gray-100 text-gray-800"} rounded-lg px-3 py-2`}
                >
                  {#if message.user_id !== data.id}
                    <div class="text-xs font-semibold mb-1">
                      {message.user_name || "Unknown User"}
                    </div>
                  {/if}
                  <p>{message.content}</p>
                  <div
                    class={`text-right text-xs ${message.user_id === data.id ? "text-purple-200" : "text-gray-500"} mt-1`}
                  >
                    {formatTime(message.timestamp)}
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
    <div class="card p-4 w-64 flex flex-col">
      <h3 class="text-base font-semibold mb-3">Team Members</h3>
      <div class="space-y-3">
        {#if !selectedRoom}
          <div class="text-center p-4 text-gray-500">
            Select a room to see members
          </div>
        {:else if teamMembers.length === 0}
          <div class="text-center p-4 text-gray-500">No team members found</div>
        {:else}
          {#each teamMembers as member}
            <div class="flex items-center">
              <!-- <img
                class="w-8 h-8 rounded-full mr-2"
                src={`/api/avatars/${member.id}`}
                alt={member.name}
              /> -->
              <div>
                <p class="text-sm font-medium">{member}</p>
                <p class="text-xs text-gray-500">
                  {member.id === data.id ? "You | " : ""}{member.role ||
                    "Member"}
                </p>
              </div>
            </div>
          {/each}
        {/if}
      </div>

      <hr class="my-3 border-gray-200" />

      <!-- Project Files -->
      <h3 class="text-base font-semibold mb-3">Shared Files</h3>
      <div class="space-y-2 flex-1 overflow-y-auto">
        {#if !selectedRoom}
          <div class="text-center p-4 text-gray-500">
            Select a room to see files
          </div>
        {:else if sharedFiles.length === 0}
          <div class="text-center p-4 text-gray-500">No shared files yet</div>
        {:else}
          {#each sharedFiles as file}
            <div class="p-2 bg-gray-50 rounded text-sm flex items-center">
              <svg
                class="w-4 h-4 text-blue-500 mr-2"
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
              <span class="truncate">{file.name}</span>
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
