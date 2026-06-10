<script lang="ts">
  import { onMount, type Snippet } from "svelte";
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";
  import { PUBLIC_WS_URL } from '$env/static/public';

  let {
    url,
    role,
    name,
    avatarUrl,
    userId,
    children,
  }: {
    url: string;
    role: string;
    name: string;
    avatarUrl: string;
    userId: string;
    children: Snippet;
  } = $props();

  // State for user dropdown menu
  let isUserMenuOpen: boolean = $state(false);
  // State for notification dropdown menu
  let isNotificationOpen: boolean = $state(false);
  // State for notification modal
  let showAllNotificationsModal: boolean = $state(false);

  // Reference to dropdown containers for click outside detection
  let userMenuContainer: HTMLDivElement | null = $state(null);
  let notificationContainer: HTMLDivElement | null = $state(null);

  // Danh sách thông báo
  let notifications: any[] = $state([]);
  let ws: WebSocket | null = null;
  let notificationReconnectTimer: ReturnType<typeof setTimeout> | null = null;
  // Invitations count
  let invitationCount: number = $state(0);

  // Toggle user dropdown
  function toggleUserMenu() {
    isUserMenuOpen = !isUserMenuOpen;
    if (isUserMenuOpen) {
      isNotificationOpen = false; // Close notification dropdown when opening user menu
    }
  }

  // Toggle notification dropdown
  function toggleNotificationMenu() {
    isNotificationOpen = !isNotificationOpen;
    if (isNotificationOpen) {
      isUserMenuOpen = false; // Close user dropdown when opening notification
    }
  }

  // Function to open all notifications modal
  function openAllNotificationsModal() {
    showAllNotificationsModal = true;
    isNotificationOpen = false; // Close dropdown when opening modal
  }

  // Function to close all notifications modal
  function closeAllNotificationsModal() {
    showAllNotificationsModal = false;
  }

  // Handle clicks outside the dropdowns to close them
  function handleClickOutside(event: MouseEvent) {
    const target = event.target as Node;

    if (userMenuContainer && !userMenuContainer.contains(target)) {
      isUserMenuOpen = false;
    }

    if (notificationContainer && !notificationContainer.contains(target)) {
      isNotificationOpen = false;
    }
  }

  onMount(() => {
    // Add global click listener when component mounts
    document.addEventListener("click", handleClickOutside);

    let shouldReconnect = true;
    if (browser && userId) {
      const connectWs = () => {
        ws = new WebSocket(`${PUBLIC_WS_URL}/ws/notifi/${userId}`);
        
        ws.onmessage = (event) => {
          try {
            const data = JSON.parse(event.data);
            if (Array.isArray(data)) {
              notifications = data.map(item => ({
                id: item.id || Date.now(),
                type: item.type || "system",
                title: item.title || "Notification",
                message: item.message || item.content || "",
                time: item.created_at ? new Date(item.created_at).toLocaleString() : new Date().toLocaleString(),
                read: item.is_read || false,
                colorClass: "border-[#896DFF]",
                bgClass: "bg-[#896DFF] bg-opacity-10",
                textClass: "text-[#896DFF]",
                icon: "M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6 6 0 00-12 0v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
              }));
            } else if (data) {
              const newNotification = {
                id: data.id || Date.now(),
                type: data.type || "system",
                title: data.title || "New Notification",
                message: data.message || data.content || "",
                time: data.created_at ? new Date(data.created_at).toLocaleString() : new Date().toLocaleString(),
                read: data.is_read || false,
                colorClass: "border-[#896DFF]",
                bgClass: "bg-[#896DFF] bg-opacity-10",
                textClass: "text-[#896DFF]",
                icon: "M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6 6 0 00-12 0v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
              };
              notifications = [newNotification, ...notifications];
            }
          } catch (e) {
            console.error("Error parsing notification:", e);
          }
        };

        ws.onclose = () => {
          if (shouldReconnect) {
            notificationReconnectTimer = setTimeout(connectWs, 5000);
          }
        };
      };
      connectWs();
    }

    // Fetch invitation count for students
    const fetchInvitations = async () => {
      try {
        const res = await fetch("/api/invitations/me");
        if (res.ok) {
          const data = await res.json();
          invitationCount = Array.isArray(data) ? data.length : 0;
        }
      } catch (e) {
        // Silent fail
      }
    };
    if (role === "student") {
      fetchInvitations();
    }

    // Clean up listener when component is destroyed
    return () => {
      shouldReconnect = false;
      document.removeEventListener("click", handleClickOutside);
      if (notificationReconnectTimer) clearTimeout(notificationReconnectTimer);
      if (ws) {
        ws.close();
      }
    };
  });

  // const logout = async () => {
  //   const response = await fetch("/api/logout", {
  //     method: "POST",
  //   });

  //   if (response.ok) {
  //     // Chuyển hướng thủ công sau khi logout thành công
  //     window.location.href = "/login";
  //   }
  // };

  // Logout client side
  const logout = async () => {
    try {
      // Gọi API server để xóa httpOnly cookie (server-side)
      await fetch("/api/logout", {
        method: "POST",
        credentials: "include",
      });

      // Xóa dữ liệu lưu trữ cục bộ
      if (browser) {
        localStorage.removeItem("user");
        localStorage.removeItem("token");
        sessionStorage.clear();
      }

      // Redirect về trang login
      window.location.href = "/login";
    } catch (error) {
      console.error("Logout failed:", error);
      window.location.href = "/login";
    }
  };

  const toggleAccountForm = () => {
    goto("/account");
  };

  const toggleSettingForm = () => {
    goto("/settings");
  };
</script>

{#if url !== "/login" && url !== "/register"}
  <aside
    class="sidebar fixed h-full w-64 p-4 flex flex-col justify-between flex-shrink-0"
  >
    <!-- Sidebar content remains unchanged -->
    <div>
      <h1
        class="text-xl font-bold mb-6 border-b border-white flex items-center"
      >
        <img src="/favicon.png" alt="SkillForge Logo" class="h-8 w-8 mr-2" />
        <span>SKILLFORGE</span>
      </h1>
      <!-- Navigation menu items -->
      <nav class="space-y-3">
        <a
          href="/dashboard"
          class={url.startsWith("/dashboard")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >
          <div class="flex items-center">
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
              ></path>
            </svg>
            Dashboard
          </div>
        </a>

        <a
          href="/project"
          class={url.startsWith("/project")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >
          <div class="flex items-center">
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"
              ></path>
            </svg>
            Projects
          </div>
        </a>

        <a
          href="/marketplace"
          class={url.startsWith("/marketplace")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >
          <div class="flex items-center">
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
              ></path>
            </svg>
            Marketplace
          </div>
        </a>

        <a
          href="/chat"
          class={url.startsWith("/chat")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >
          <div class="flex items-center">
            <svg
              class="w-5 h-5 mr-2"
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
              ></path>
            </svg>
            Chats
          </div>
        </a>

        <a
          href="/profile"
          class={url.startsWith("/profile")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >
          <div class="flex items-center">
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
              ></path>
            </svg>
            Profile
          </div>
        </a>

        {#if role === "student"}
          <a
            href="/invitations"
            class={url.startsWith("/invitations")
              ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
              : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
                Invitations
              </div>
              {#if invitationCount > 0}
                <span class="bg-[#ff6f61] text-white text-xs font-bold px-2 py-0.5 rounded-full">{invitationCount}</span>
              {/if}
            </div>
          </a>
        {/if}
      </nav>
    </div>

    <!-- User profile section -->
    <div class="mt-auto pt-4 border-t border-white">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <img
            class="w-10 h-10 rounded-full border-2 border-white"
            src={avatarUrl}
            alt="User Avatar"
          />
          <div>
            <p class="text-sm font-medium">{name}</p>
            <p class="text-xs text-gray-400">
              {role.charAt(0).toUpperCase() + role.slice(1)}
            </p>
          </div>
        </div>

        <div class="flex space-x-2">
          <!-- Notification Icon -->
          <div class="relative" bind:this={notificationContainer}>
            <button
              aria-label="Notifications"
              class="p-1 rounded-full hover:bg-[#896DFF] focus:outline-none"
              onclick={toggleNotificationMenu}
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
                  d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6 6 0 00-12 0v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                ></path>
              </svg>
              <span
                class="absolute top-0 right-0 w-2 h-2 bg-[#ff6f61] rounded-full"
              ></span>
            </button>
          </div>

          <!-- Logout Icon -->
          <button
            aria-label="Logout"
            onclick={logout}
            class="p-1 rounded-full hover:bg-[#896DFF] focus:outline-none"
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
                d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
              ></path>
            </svg>
          </button>
        </div>
      </div>
    </div>
  </aside>

  <!-- All Notifications Modal -->
  {#if showAllNotificationsModal}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center modal"
    >
      <div
        class="bg-white rounded-lg shadow-lg w-full max-w-3xl max-h-[90vh] flex flex-col"
      >
        <!-- Modal Header -->
        <div
          class="flex justify-between items-center px-6 py-4 border-b border-gray-200"
        >
          <h2 class="text-xl font-semibold text-gray-800">All Notifications</h2>
          <button
            aria-label="Close notifications"
            class="text-gray-500 hover:text-gray-700"
            onclick={closeAllNotificationsModal}
          >
            <svg
              class="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              ></path>
            </svg>
          </button>
        </div>

        <!-- Notification List - Scrollable Content -->
        <div class="overflow-y-auto flex-grow">
          {#each notifications as notification}
            <div
              class="p-4 border-b border-gray-100 hover:bg-gray-50 cursor-pointer flex"
              class:bg-gray-50={!notification.read}
            >
              <div class="flex items-start w-full">
                <div
                  class="{notification.bgClass} p-2 rounded-full mr-3 flex-shrink-0"
                >
                  <svg
                    class="w-5 h-5 {notification.textClass}"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d={notification.icon}
                    ></path>
                  </svg>
                </div>
                <div class="flex-1">
                  <div class="flex justify-between items-start">
                    <p class="font-semibold text-gray-800">
                      {notification.title}
                      {#if !notification.read}
                        <span
                          class="ml-2 w-2 h-2 bg-[#ff6f61] rounded-full inline-block"
                        ></span>
                      {/if}
                    </p>
                    <p class="text-xs text-gray-500">{notification.time}</p>
                  </div>
                  <p class="text-sm text-gray-600 mt-1">
                    {notification.message}
                  </p>

                  <!-- Action buttons -->
                  <div class="mt-2 flex gap-2">
                    <button class="text-xs text-blue-600 hover:underline"
                      >View details</button
                    >
                    <button class="text-xs text-gray-500 hover:underline"
                      >Mark as {notification.read ? "unread" : "read"}</button
                    >
                  </div>
                </div>
              </div>
            </div>
          {/each}
        </div>

        <!-- Modal Footer -->
        <div
          class="px-6 py-3 border-t border-gray-200 bg-gray-50 flex justify-end"
        >
          <button
            class="px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300"
            onclick={closeAllNotificationsModal}
          >
            Close
          </button>
        </div>
      </div>
    </div>
  {/if}
{/if}

{#if isNotificationOpen}
  <div class="fixed top-0 left-0 w-full h-full z-[9999] pointer-events-none">
    <div
      class="absolute bottom-15 left-30 mb-2 w-100 bg-white border border-gray-200 rounded-lg shadow-lg pointer-events-auto"
    >
      <div
        class="px-4 py-3 border-b border-gray-200 flex justify-between items-center"
      >
        <h3 class="font-semibold text-gray-700">Notifications</h3>
      </div>
      <div class="max-h-72 overflow-y-auto">
        {#each notifications.slice(0, 3) as notification}
          <div
            class="p-3 border-l-4 {notification.colorClass} hover:bg-gray-50 cursor-pointer"
          >
            <div class="flex items-start">
              <div class="{notification.bgClass} p-2 rounded-full mr-3">
                <svg
                  class="w-5 h-5 {notification.textClass}"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d={notification.icon}
                  ></path>
                </svg>
              </div>
              <div class="flex-1">
                <p class="font-semibold text-sm text-gray-800">
                  {notification.title}
                </p>
                <p class="text-xs text-gray-600 mt-1">
                  {notification.message}
                </p>
                <p class="text-xs text-gray-500 mt-2">
                  {notification.time}
                </p>
              </div>
            </div>
          </div>
        {/each}
      </div>

      <div class="px-4 py-3 text-center border-t border-gray-200">
        <button
          class="text-sm text-[#896DFF] font-medium hover:underline w-full"
          onclick={openAllNotificationsModal}
        >
          View All Notifications
        </button>
      </div>
    </div>
  </div>
{/if}

{@render children()}
