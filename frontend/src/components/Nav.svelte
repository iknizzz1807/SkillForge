<script lang="ts">
  import { onMount, type Snippet } from "svelte";
  let {
    url,
    role,
    name,
    children,
  }: {
    url: string;
    role: string;
    name: string;
    children: Snippet;
  } = $props();

  // State for user dropdown menu
  let isUserMenuOpen: boolean = $state(false);
  // State for notification dropdown menu
  let isNotificationOpen: boolean = $state(false);

  // Reference to dropdown containers for click outside detection
  let userMenuContainer: HTMLDivElement | null = $state(null);
  let notificationContainer: HTMLDivElement | null = $state(null);

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

    // Clean up listener when component is destroyed
    return () => {
      document.removeEventListener("click", handleClickOutside);
    };
  });

  const logout = async () => {
    const response = await fetch("/api/logout", {
      method: "POST",
    });

    if (response.ok) {
      // Chuyển hướng thủ công sau khi logout thành công
      window.location.href = "/login";
    }
  };

  const toggleAccountForm = () => {
    //
  };

  const toggleSettingForm = () => {
    //
  };
</script>

{#if url !== "/login" && url !== "/register"}
  <aside
    class="sidebar fixed h-full w-64 p-4 flex flex-col justify-between flex-shrink-0"
  >
    <div>
      <h1 class="text-xl font-bold mb-6 border-b border-white">SKILLFORGE</h1>
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
      </nav>
    </div>

    <!-- User profile section remains unchanged -->
    <div class="mt-auto pt-4 border-t border-white">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <img
            class="w-10 h-10 rounded-full border-2 border-white"
            src="https://avatars.githubusercontent.com/u/123456?v=4"
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

            <div
              class={`absolute bottom-full left-0 mb-2 w-64 bg-white border border-gray-200 rounded shadow-lg ${!isNotificationOpen ? "hidden" : ""}`}
            >
              <!-- 2 notifications -->
              <div
                class="p-2 text-sm text-gray-800 hover:bg-gray-100 cursor-pointer"
              >
                <p class="font-semibold">Task submitted!</p>
                <span class="text-xs text-gray-500">18/03/2025</span>
              </div>

              <div
                class="p-2 text-sm text-gray-800 border-t border-gray-200 hover:bg-gray-100 cursor-pointer"
              >
                <p class="font-semibold">Project updated</p>
                <span class="text-xs text-gray-500">17/03/2025</span>
              </div>
            </div>
          </div>

          <!-- Logout Icon -->
          <button
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

  <!-- <header class="flex justify-between items-center mb-4 ml-64 pr-4 pl-4 pt-4">
    <div>
      <h2 class="text-xl font-semibold">{header}</h2>
      <p class="text-sm text-gray-600">{description}</p>
    </div>
  </header> -->
{/if}

{@render children()}
