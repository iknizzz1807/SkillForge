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

  let isSidebarOpen = $state(false);
  let isNotificationOpen = $state(false);
  let showAllNotificationsModal = $state(false);

  let notificationContainer: HTMLDivElement | null = $state(null);

  let notifications: any[] = $state([]);
  let ws: WebSocket | null = null;
  let notificationReconnectTimer: ReturnType<typeof setTimeout> | null = null;
  let invitationCount: number = $state(0);

  function closeSidebar() { isSidebarOpen = false; }

  function toggleNotificationMenu() {
    isNotificationOpen = !isNotificationOpen;
  }

  function openAllNotificationsModal() {
    showAllNotificationsModal = true;
    isNotificationOpen = false;
  }

  function closeAllNotificationsModal() {
    showAllNotificationsModal = false;
  }

  function handleClickOutside(event: MouseEvent) {
    if (notificationContainer && !notificationContainer.contains(event.target as Node)) {
      isNotificationOpen = false;
    }
  }

  onMount(() => {
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
              }));
            } else if (data) {
              notifications = [{
                id: data.id || Date.now(),
                type: data.type || "system",
                title: data.title || "New Notification",
                message: data.message || data.content || "",
                time: data.created_at ? new Date(data.created_at).toLocaleString() : new Date().toLocaleString(),
                read: data.is_read || false,
              }, ...notifications];
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

    const fetchInvitations = async () => {
      try {
        const res = await fetch("/api/invitations/me");
        if (res.ok) {
          const data = await res.json();
          invitationCount = Array.isArray(data) ? data.length : 0;
        }
      } catch (e) { /* silent */ }
    };
    if (role === "student") {
      fetchInvitations();
    }

    return () => {
      shouldReconnect = false;
      document.removeEventListener("click", handleClickOutside);
      if (notificationReconnectTimer) clearTimeout(notificationReconnectTimer);
      if (ws) ws.close();
    };
  });

  const logout = async () => {
    try {
      await fetch("/api/logout", { method: "POST", credentials: "include" });
      if (browser) {
        localStorage.removeItem("user");
        localStorage.removeItem("token");
        sessionStorage.clear();
      }
      window.location.href = "/login";
    } catch {
      window.location.href = "/login";
    }
  };

  const navItems = [
    { href: "/dashboard", label: "Dashboard", icon: "M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" },
    { href: "/project", label: "Projects", icon: "M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" },
    { href: "/marketplace", label: "Marketplace", icon: "M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" },
    { href: "/chat", label: "Chats", icon: "M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" },
    { href: "/profile", label: "Profile", icon: "M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" },
  ];
</script>

{#if url !== "/login" && url !== "/register"}
  <!-- Mobile header -->
  <div class="md:hidden fixed top-0 left-0 right-0 h-16 bg-gray-950 border-b border-gray-800 flex items-center justify-between px-4 z-40">
    <h1 class="text-lg font-bold text-white flex items-center gap-2">
      <img src="/favicon.png" alt="SkillForge" class="h-7 w-7" />
      <span>SKILLFORGE</span>
    </h1>
    <button class="p-2 text-gray-400 hover:text-white" onclick={() => (isSidebarOpen = true)} aria-label="Open menu">
      <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
      </svg>
    </button>
  </div>

  <!-- Backdrop -->
  {#if isSidebarOpen}
    <div class="fixed inset-0 bg-black/60 z-40 md:hidden" onclick={closeSidebar}></div>
  {/if}

  <!-- Sidebar -->
  <aside
    class="fixed top-0 left-0 h-full w-64 bg-gray-950 text-white border-r border-gray-800 flex flex-col z-50 transition-transform duration-300 md:translate-x-0 {isSidebarOpen ? 'translate-x-0' : '-translate-x-full'}"
  >
    <!-- Logo -->
    <div class="p-4 border-b border-gray-800 hidden md:block">
      <h1 class="text-xl font-bold flex items-center gap-2">
        <img src="/favicon.png" alt="SkillForge" class="h-8 w-8" />
        <span>SKILLFORGE</span>
      </h1>
    </div>

    <!-- Nav items -->
    <nav class="flex-1 p-3 space-y-1 overflow-y-auto">
      {#each navItems as item}
        <a
          href={item.href}
          onclick={closeSidebar}
          class={"flex items-center gap-3 px-3 py-2 rounded-lg text-sm transition-colors " + (url.startsWith(item.href) ? "bg-[#896DFF] text-white font-semibold" : "text-gray-300 hover:text-white hover:bg-[#896DFF]/10")}
        >
          <svg class="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d={item.icon} />
          </svg>
          {item.label}
        </a>
      {/each}

      {#if role === "student"}
        <a
          href="/invitations"
          onclick={closeSidebar}
          class={"flex items-center justify-between px-3 py-2 rounded-lg text-sm transition-colors " + (url.startsWith("/invitations") ? "bg-[#896DFF] text-white font-semibold" : "text-gray-300 hover:text-white hover:bg-[#896DFF]/10")}
        >
          <div class="flex items-center gap-3">
            <svg class="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
            Invitations
          </div>
          {#if invitationCount > 0}
            <span class="bg-[#ff6f61] text-white text-xs font-bold px-2 py-0.5 rounded-full">{invitationCount}</span>
          {/if}
        </a>
      {/if}
    </nav>

    <!-- User section -->
    <div class="p-3 border-t border-gray-800">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3 min-w-0 flex-1">
          <img class="w-9 h-9 rounded-full object-cover border border-gray-700 flex-shrink-0" src={avatarUrl} alt="" />
          <div class="min-w-0">
            <p class="text-sm font-medium truncate">{name}</p>
            <p class="text-xs text-gray-400 truncate">{role.charAt(0).toUpperCase() + role.slice(1)}</p>
          </div>
        </div>
        <div class="flex gap-1 flex-shrink-0">
          <div class="relative" bind:this={notificationContainer}>
            <button class="p-1.5 rounded-full hover:bg-gray-800 text-gray-400 hover:text-white" onclick={toggleNotificationMenu} aria-label="Notifications">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6 6 0 00-12 0v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
              </svg>
              {#if notifications.some(n => !n.read)}
                <span class="absolute top-1 right-1 w-2.5 h-2.5 bg-[#ff6f61] rounded-full"></span>
              {/if}
            </button>
          </div>
          <button class="p-1.5 rounded-full hover:bg-gray-800 text-gray-400 hover:text-white" onclick={logout} aria-label="Logout">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </aside>

  <!-- All Notifications Modal -->
  {#if showAllNotificationsModal}
    <div class="fixed inset-0 bg-black/60 z-[100] flex items-center justify-center p-4">
      <div class="bg-white rounded-xl shadow-xl w-full max-w-2xl max-h-[85vh] flex flex-col overflow-hidden">
        <div class="flex items-center justify-between px-6 py-4 border-b border-gray-100 flex-shrink-0">
          <h2 class="text-lg font-bold text-gray-900">All Notifications</h2>
          <button class="text-gray-400 hover:text-gray-600 rounded-lg p-1.5 hover:bg-gray-100" onclick={closeAllNotificationsModal} aria-label="Close">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <div class="overflow-y-auto flex-1 min-h-0 divide-y divide-gray-50">
          {#each notifications as notification}
            <div class={"p-4 flex gap-3 " + (!notification.read ? "bg-purple-50/30" : "")}>
              <div class="bg-[#896DFF]/10 p-2 rounded-xl flex-shrink-0 h-9 w-9 flex items-center justify-center">
                <svg class="w-5 h-5 text-[#896DFF]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6 6 0 00-12 0v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                </svg>
              </div>
              <div class="flex-1 min-w-0">
                <div class="flex justify-between items-start mb-1">
                  <p class="font-semibold text-sm text-gray-900">{notification.title}</p>
                  <p class="text-[10px] text-gray-400 flex-shrink-0">{notification.time}</p>
                </div>
                <p class="text-xs text-gray-600">{notification.message}</p>
              </div>
            </div>
          {/each}
        </div>
        <div class="px-6 py-3 border-t border-gray-100 bg-gray-50 flex justify-end flex-shrink-0">
          <button class="btn-secondary text-sm px-4 py-2" onclick={closeAllNotificationsModal}>Close</button>
        </div>
      </div>
    </div>
  {/if}
{/if}

<!-- Notification dropdown -->
{#if isNotificationOpen}
  <div class="fixed inset-0 z-50 pointer-events-none">
    <div class="absolute bottom-16 left-4 md:left-64 mb-2 w-80 bg-white border border-gray-200 rounded-xl shadow-xl pointer-events-auto overflow-hidden">
      <div class="px-4 py-3 border-b border-gray-100 bg-gray-50">
        <h3 class="font-bold text-xs text-gray-700">Notifications</h3>
      </div>
      <div class="max-h-64 overflow-y-auto divide-y divide-gray-50">
        {#each notifications.slice(0, 3) as notification}
          <div class="p-3 hover:bg-gray-50 flex gap-2">
            <div class="bg-[#896DFF]/10 p-1.5 rounded-lg flex-shrink-0 h-8 w-8 flex items-center justify-center">
              <svg class="w-4 h-4 text-[#896DFF]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6 6 0 00-12 0v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
              </svg>
            </div>
            <div class="flex-1 min-w-0">
              <p class="font-semibold text-xs text-gray-900 truncate">{notification.title}</p>
              <p class="text-[11px] text-gray-500 truncate">{notification.message}</p>
            </div>
          </div>
        {/each}
      </div>
      <div class="px-4 py-2 text-center border-t border-gray-100 bg-gray-50">
        <button class="text-xs text-[#896DFF] font-semibold hover:underline w-full" onclick={openAllNotificationsModal}>
          View All Notifications
        </button>
      </div>
    </div>
  </div>
{/if}

<!-- Main content wrapper -->
<div class="flex-1 w-full md:pl-64 pt-16 md:pt-0 min-h-screen">
  {@render children()}
</div>
