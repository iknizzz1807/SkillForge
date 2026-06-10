import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/svelte";
import MarketplacePage from "../marketplace/+page.svelte";

describe("Marketplace page", () => {
  const mockProps = {
    data: {
      role: "student",
      projects: [
        {
          id: "1",
          title: "Build a Chat App",
          description: "Cool project",
          skills: ["React", "Node.js"],
          start_time: "2025-03-01",
          end_time: "2025-04-01",
          created_at: "2025-02-15",
        },
        {
          id: "2",
          title: "E-commerce Backend",
          description: "Backend project",
          skills: ["Python"],
          start_time: "2025-03-10",
          end_time: "2025-04-10",
          created_at: "2025-02-20",
        },
      ],
    },
  };

  it("renders project count heading", () => {
    render(MarketplacePage, { props: mockProps });
    expect(screen.getByText(/Projects\s*\(2\)/)).toBeInTheDocument();
  });

  it("renders project cards with titles", () => {
    render(MarketplacePage, { props: mockProps });
    expect(screen.getByText("Build a Chat App")).toBeInTheDocument();
    expect(screen.getByText("E-commerce Backend")).toBeInTheDocument();
  });

  it("renders skills as tags", () => {
    render(MarketplacePage, { props: mockProps });
    expect(screen.getByText("React")).toBeInTheDocument();
    expect(screen.getByText("Python")).toBeInTheDocument();
  });

  it("shows no projects message when empty", () => {
    render(MarketplacePage, {
      props: { data: { role: "student", projects: [] } },
    });
    expect(screen.getByText(/No projects available/i)).toBeInTheDocument();
  });
});
