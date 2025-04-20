<script lang="ts">
  import "../app.css";
  let { children, data }: { children: Snippet; data: LayoutData } = $props();
  import Nav from "../components/Nav.svelte";
  import { page } from "$app/state";
  import type { LayoutData } from "./$types"; // để lấy giữ liệu người dùng hiện tại mà hiển thị cho chuẩn xác
  import type { Snippet } from "svelte";

  let header: string = $state("");
  let description: string = $state("");

  let role: string = $state("");
  let name: string = $state("");
  let avatarUrl: string = $state("");

  $effect(() => {
    if (data.role && data.userName && data.avatarUrl) {
      role = data.role;
      name = data.userName;
      avatarUrl = data.avatarUrl;
    }
  });

  $effect(() => {
    if (page.url.pathname) {
      if (page.url.pathname.startsWith("/project")) {
        header = "Your projects";
        description = "Manage your projects";
      } else if (page.url.pathname.startsWith("/dashboard")) {
        header = "Welcome back, " + data.userName;
        description = new Date().toLocaleString();
      } else if (page.url.pathname.startsWith("/marketplace")) {
        header = "Project Marketplace";
        description = "Find and apply to projects that match your skills";
      } else if (page.url.pathname.startsWith("/chat")) {
        header = "Chat";
        description = "Chat with your collaborators and business";
      } else if (page.url.pathname.startsWith("/profile")) {
        header = "Your Profile";
        description = "Manage your personal information";
      } else if (page.url.pathname.startsWith("/analytics")) {
        header = "Analytics";
        description = "Track your performance and skills";
      }
    }
  });
</script>

<svelte:head>
  <title>{header}</title>
</svelte:head>

<Nav url={page.url.pathname} {role} {name} {avatarUrl}>
  {@render children()}
</Nav>
