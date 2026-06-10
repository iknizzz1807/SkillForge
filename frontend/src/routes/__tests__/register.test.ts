import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/svelte";
import RegisterPage from "../register/+page.svelte";

const registerProps = {
  data: {
    form: null,
    data: { role: "" },
  },
  form: undefined,
};

describe("Register page", () => {
  it("renders the Register heading", () => {
    render(RegisterPage, { props: registerProps });
    expect(screen.getByText("Register to SKILLFORGE")).toBeInTheDocument();
  });

  it("renders Full Name input field", () => {
    render(RegisterPage, { props: registerProps });
    const input = screen.getByLabelText(/Full Name/);
    expect(input).toBeInTheDocument();
  });

  it("renders Email input field", () => {
    render(RegisterPage, { props: registerProps });
    const input = screen.getByLabelText(/Email/);
    expect(input).toBeInTheDocument();
  });

  it("renders Password input field", () => {
    render(RegisterPage, { props: registerProps });
    const input = screen.getAllByLabelText(/Password/i)[0];
    expect(input).toBeInTheDocument();
  });

  it("renders role selection (Student / Business)", () => {
    render(RegisterPage, { props: registerProps });
    expect(screen.getByText("Student")).toBeInTheDocument();
    expect(screen.getByText("Business")).toBeInTheDocument();
  });

  it("shows error message when form has error", () => {
    render(RegisterPage, {
      props: {
        ...registerProps,
        form: { error: true, message: "Email already exists" },
      },
    });
    expect(screen.getByText("Email already exists")).toBeInTheDocument();
  });

  it("renders a link to login page", () => {
    render(RegisterPage, { props: registerProps });
    const loginLink = screen.getByRole("link", { name: /login/i });
    expect(loginLink).toBeInTheDocument();
    expect(loginLink).toHaveAttribute("href", "/login");
  });

  it("renders Create Account submit button", () => {
    render(RegisterPage, { props: registerProps });
    const btn = screen.getByRole("button", { name: /Create Account/i });
    expect(btn).toBeInTheDocument();
  });

  it("renders avatar upload area", () => {
    render(RegisterPage, { props: registerProps });
    const label = screen.getByLabelText("Upload avatar");
    expect(label).toBeInTheDocument();
  });
});
