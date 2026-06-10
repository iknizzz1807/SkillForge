import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/svelte";
import ChatPage from "../+page.svelte";

const mockData = {
  data: { id: "user1", userName: "Alice", avatarUrl: "/avatars/user1" },
};

describe("Chat page", () => {
  it("renders the Conversations heading", () => {
    render(ChatPage, { props: mockData });
    expect(screen.getByText("Conversations")).toBeInTheDocument();
  });

  it("renders search input for conversations", () => {
    render(ChatPage, { props: mockData });
    const searchInput = screen.getByPlaceholderText("Search conversations...");
    expect(searchInput).toBeInTheDocument();
  });

  it("shows no conversations message when list is empty", () => {
    render(ChatPage, { props: mockData });
    expect(screen.getByText("No conversations found")).toBeInTheDocument();
  });

  it("shows select a conversation message initially", () => {
    render(ChatPage, { props: mockData });
    expect(screen.getByText("Select a conversation")).toBeInTheDocument();
  });

  it("shows no team selected message when no room selected", () => {
    render(ChatPage, { props: mockData });
    expect(screen.getByText("No team selected")).toBeInTheDocument();
  });

  it("renders empty chat area description", () => {
    render(ChatPage, { props: mockData });
    expect(
      screen.getByText("Choose a chat room to start messaging")
    ).toBeInTheDocument();
  });

  it("shows conversation count badge", () => {
    render(ChatPage, { props: mockData });
    const badge = screen.getByText("0");
    expect(badge).toBeInTheDocument();
  });
});
