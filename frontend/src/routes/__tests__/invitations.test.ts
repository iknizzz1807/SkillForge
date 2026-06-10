import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/svelte";
import InvitationsPage from "../invitations/+page.svelte";

describe("Invitations page", () => {
  it("renders the page heading", () => {
    render(InvitationsPage, { props: { data: {} } });
    expect(screen.getByText("My Invitations")).toBeInTheDocument();
  });

  it("renders page description", () => {
    render(InvitationsPage, { props: { data: {} } });
    expect(screen.getByText("Manage your project invitations")).toBeInTheDocument();
  });

  it("shows loading spinner initially", () => {
    render(InvitationsPage, { props: { data: {} } });
    const spinner = document.querySelector(".animate-spin");
    expect(spinner).toBeInTheDocument();
  });
});
