import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/svelte";
import DashboardPage from "../dashboard/+page.svelte";

describe("Dashboard page (student)", () => {
  const studentProps = {
    data: { role: "student", user: { id: "1", name: "Alice", role: "student" } },
  };

  it("renders Student Dashboard heading", () => {
    render(DashboardPage, { props: studentProps });
    expect(screen.getByText("Student Dashboard")).toBeInTheDocument();
  });

  it("renders Active Projects stat", () => {
    render(DashboardPage, { props: studentProps });
    const statHeadings = screen.getAllByText("Active Projects");
    expect(statHeadings.length).toBeGreaterThan(0);
  });

  it("renders suggested projects section", () => {
    render(DashboardPage, { props: studentProps });
    const suggested = screen.getByText(/Suggested Projects/i);
    expect(suggested).toBeInTheDocument();
  });

  it("renders Apply buttons on suggested projects", () => {
    render(DashboardPage, { props: studentProps });
    const applyButtons = screen.getAllByText("Apply");
    expect(applyButtons.length).toBeGreaterThan(0);
  });

  it("renders student-specific charts section", () => {
    render(DashboardPage, { props: studentProps });
    expect(screen.getByText("Your Skills Assessment")).toBeInTheDocument();
  });
});

describe("Dashboard page (business)", () => {
  const businessProps = {
    data: { role: "business", user: { id: "2", name: "Biz", role: "business" } },
  };

  it("renders Business Dashboard heading", () => {
    render(DashboardPage, { props: businessProps });
    expect(screen.getByText("Business Dashboard")).toBeInTheDocument();
  });

  it("renders Create Project link", () => {
    render(DashboardPage, { props: businessProps });
    const createProject = screen.getByText("Create Project");
    expect(createProject).toBeInTheDocument();
  });

  it("renders Find Talent link", () => {
    render(DashboardPage, { props: businessProps });
    expect(screen.getByText("Find Talent")).toBeInTheDocument();
  });

  it("renders Messages link", () => {
    render(DashboardPage, { props: businessProps });
    expect(screen.getByText("Messages")).toBeInTheDocument();
  });

  it("renders Projects Status Overview", () => {
    render(DashboardPage, { props: businessProps });
    expect(screen.getByText("Projects Status Overview")).toBeInTheDocument();
  });
});
