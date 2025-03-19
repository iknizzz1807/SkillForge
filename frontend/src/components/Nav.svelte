<script lang="ts">
  import { onMount, type Snippet } from "svelte";

  let {
    header,
    description,
    url,
    children,
  }: { header: string; description: string; url: string; children: Snippet } =
    $props();

  // console.log(url);

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
    await fetch("/api/logout", {
      method: "POST",
    });
  };

  const toggleAccountForm = () => {
    //
  };

  const toggleSettingForm = () => {
    //
  };
</script>

{#if url !== "/login" && url !== "/register"}
  <aside class="sidebar fixed h-full w-56 p-4 flex-shrink-0">
    <h1 class="text-xl font-bold mb-6">SKILLFORGE</h1>
    <nav class="space-y-3">
      <a
        href="/dashboard"
        class={url === "dashboard"
          ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
          : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >Dashboard</a
      >
      <a
        href="/project"
        class={url === "project"
          ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
          : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >Projects</a
      >
      <a
        href="/marketplace"
        class={url === "marketplace"
          ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
          : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >Marketplace</a
      >
      <a
        href="/chat"
        class={url === "chat"
          ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
          : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >Chats</a
      >
      <a
        href="/profile"
        class={url === "profile"
          ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
          : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >Profile</a
      >
      <a
        href="/analytics"
        class={url === "analytics"
          ? "block text-base font-bold bg-[#896DFF] rounded px-2 py-1"
          : "block text-base hover:bg-[#896DFF] hover:bg-opacity-20 rounded px-2 py-1"}
        >Analytics</a
      >
    </nav>
  </aside>

  <header class="flex justify-between items-center mb-4 ml-56 p-4">
    <div>
      <h2 class="text-xl font-semibold">{header}</h2>
      <p class="text-sm text-gray-600">{description}</p>
    </div>
    <div class="flex items-center space-x-3">
      <!-- Notification Icon -->
      <!-- svelte-ignore a11y_consider_explicit_label a11y_click_events_have_key_events, a11y_no_static_element_interactions (because of reasons) -->
      <div class="relative" bind:this={notificationContainer}>
        <button class="focus:outline-none" onclick={toggleNotificationMenu}>
          <svg
            class="w-6 h-6 text-gray-600 hover:text-[#6b48ff]"
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
          <span class="absolute top-0 right-0 w-2 h-2 bg-[#ff6f61] rounded-full"
          ></span>
        </button>
        <div
          class={`absolute right-0 mt-2 w-64 bg-white border border-gray-200 rounded shadow-lg ${!isNotificationOpen ? "hidden" : ""}`}
        >
          <div class="p-2 text-sm">
            <p class="font-semibold">Task submitted!</p>
            <span class="text-xs text-gray-500">18/03/2025</span>
          </div>
        </div>
      </div>
      <!-- User Avatar -->
      <div class="relative" bind:this={userMenuContainer}>
        <button class="focus:outline-none" onclick={toggleUserMenu}>
          <img
            class="w-8 h-8 rounded-full"
            src="https://avatars.githubusercontent.com/u/123456?v=4"
            alt="User Avatar"
          />
        </button>
        <div
          class={`absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded shadow-lg ${!isUserMenuOpen ? "hidden" : ""}`}
        >
          <button
            class="block px-4 py-2 text-sm hover:bg-gray-100 w-full text-left"
            onclick={toggleAccountForm}>Account</button
          >
          <button
            class="block px-4 py-2 text-sm hover:bg-gray-100 w-full text-left"
            onclick={toggleSettingForm}>Settings</button
          >
          <button
            class="block px-4 py-2 text-sm hover:bg-gray-100 w-full text-left"
            onclick={logout}>Log Out</button
          >
        </div>
      </div>
    </div>
  </header>
{/if}

{@render children()}
