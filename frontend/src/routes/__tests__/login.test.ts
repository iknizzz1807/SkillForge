import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/svelte";
import LoginPage from "../login/+page.svelte";

const loginProps = {
  data: { redirectTo: "/dashboard" },
  form: undefined,
};

describe("Login page", () => {
  it("renders Login heading", () => {
    render(LoginPage, { props: loginProps });
    expect(screen.getByText("Login to SKILLFORGE")).toBeInTheDocument();
  });

  it("renders email input field", () => {
    render(LoginPage, { props: loginProps });
    const input = screen.getByLabelText(/Email/);
    expect(input).toBeInTheDocument();
  });

  it("renders password input field", () => {
    render(LoginPage, { props: loginProps });
    const input = screen.getByLabelText(/Password/);
    expect(input).toBeInTheDocument();
  });

  it("renders a submit button", () => {
    render(LoginPage, { props: loginProps });
    const btn = screen.getByRole("button", { name: /login/i });
    expect(btn).toBeInTheDocument();
  });

  it("has email input as required", () => {
    render(LoginPage, { props: loginProps });
    const input = screen.getByLabelText(/Email/);
    expect(input).toHaveAttribute("required");
  });

  it("has password input as required", () => {
    render(LoginPage, { props: loginProps });
    const input = screen.getByLabelText(/Password/);
    expect(input).toHaveAttribute("required");
  });

  it("renders a link to register page", () => {
    render(LoginPage, { props: loginProps });
    const registerLink = screen.getByRole("link", { name: /sign up/i });
    expect(registerLink).toBeInTheDocument();
    expect(registerLink).toHaveAttribute("href", "/register");
  });

  it("shows error message when form has error", () => {
    render(LoginPage, {
      props: { ...loginProps, form: { error: true } },
    });
    expect(screen.getByText("Invalid credentials")).toBeInTheDocument();
  });

  it("has hidden redirectTo input", () => {
    render(LoginPage, { props: loginProps });
    const input = document.querySelector('input[name="redirectTo"]');
    expect(input).toBeInTheDocument();
    expect(input).toHaveAttribute("type", "hidden");
  });
});
