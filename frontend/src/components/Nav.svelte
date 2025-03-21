<script lang="ts">
  import { onMount, type Snippet } from "svelte";
  import type { PageData } from "../routes/$types";

  let {
    header,
    description,
    url,
    role,
    name,
    children,
  }: {
    header: string;
    description: string;
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
          >Dashboard</a
        >
        <a
          href="/project"
          class={url.startsWith("/project")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
          >Projects</a
        >
        <a
          href="/marketplace"
          class={url.startsWith("/marketplace")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
          >Marketplace</a
        >
        <a
          href="/chat"
          class={url.startsWith("/chat")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
          >Chats</a
        >
        <a
          href="/profile"
          class={url.startsWith("/profile")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
          >Profile</a
        >
        <a
          href="/analytics"
          class={url.startsWith("/analytics")
            ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
            : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
          >Analytics</a
        >
      </nav>
    </div>

    <!-- User profile section at bottom of sidebar -->
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
            <p class="text-xs text-gray-400">{role}</p>
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
                class="w-5 h-5 text-gray-300 hover:text-white"
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

          <!-- Profile Icon -->
          <!-- <a
            href="/profile"
            class={url.startsWith("/profile")
              ? "p-1 rounded-full hover:bg-[#896DFF] focus:outline-none bg-[#896DFF]"
              : "p-1 rounded-full hover:bg-[#896DFF] focus:outline-none"}
          >
            <svg
              class="w-5 h-5 text-gray-300 hover:text-white"
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
          </a> -->

          <!-- Logout Icon -->
          <button
            onclick={logout}
            class="p-1 rounded-full hover:bg-[#896DFF] focus:outline-none"
          >
            <svg
              class="w-5 h-5 text-gray-300 hover:text-white"
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

  <header class="flex justify-between items-center mb-4 ml-64 pr-4 pl-4 pt-4">
    <div>
      <h2 class="text-xl font-semibold">{header}</h2>
      <p class="text-sm text-gray-600">{description}</p>
    </div>
  </header>
{/if}

{@render children()}
