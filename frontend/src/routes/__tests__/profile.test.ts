import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/svelte";
import ProfilePage from "../profile/+page.svelte";

const studentProps = {
  data: {
    id: "user1",
    name: "Alice",
    email: "alice@example.com",
    role: "student",
    title: "Software Engineer",
    avatarUrl: "/avatars/user1",
    userData: {},
    businessInfo: { companyType: "", founded: "", companySize: "", website: "", aboutUs: "" },
  },
};

describe("Profile page (student)", () => {
  it("renders Student Profile heading", () => {
    render(ProfilePage, { props: studentProps });
    expect(screen.getByText("Student Profile")).toBeInTheDocument();
  });

  it("renders user name", () => {
    render(ProfilePage, { props: studentProps });
    expect(screen.getByText("Alice")).toBeInTheDocument();
  });

  it("renders user email", () => {
    render(ProfilePage, { props: studentProps });
    expect(screen.getByText("alice@example.com")).toBeInTheDocument();
  });

  it("renders Edit Profile button", () => {
    render(ProfilePage, { props: studentProps });
    expect(screen.getByText("Edit Profile")).toBeInTheDocument();
  });

  it("renders Skills & Expertise section for student", () => {
    render(ProfilePage, { props: studentProps });
    expect(screen.getByText("Skills & Expertise")).toBeInTheDocument();
  });

  it("renders Portfolio Generator for student", () => {
    render(ProfilePage, { props: studentProps });
    expect(screen.getByText("Generate Portfolio")).toBeInTheDocument();
  });

  it("renders Badges & Achievements section", () => {
    render(ProfilePage, { props: studentProps });
    expect(screen.getByText("Badges & Achievements")).toBeInTheDocument();
  });

  it("renders XP bar for student", () => {
    render(ProfilePage, { props: studentProps });
    expect(screen.getByText("1250 XP")).toBeInTheDocument();
  });
});

const businessProps = {
  data: {
    id: "user2",
    name: "BizCorp",
    email: "biz@example.com",
    role: "business",
    title: "CTO",
    avatarUrl: "/avatars/user2",
    userData: {},
    businessInfo: {
      companyType: "Tech",
      founded: "2020",
      companySize: "50-100",
      website: "bizcorp.com",
      aboutUs: "We build stuff",
    },
  },
};

describe("Profile page (business)", () => {
  it("renders Business Profile heading", () => {
    render(ProfilePage, { props: businessProps });
    expect(screen.getByText("Business Profile")).toBeInTheDocument();
  });

  it("renders Company Information section", () => {
    render(ProfilePage, { props: businessProps });
    expect(screen.getByText("Company Information")).toBeInTheDocument();
  });

  it("renders Engagement Statistics", () => {
    render(ProfilePage, { props: businessProps });
    expect(screen.getByText("Engagement Statistics")).toBeInTheDocument();
  });

  it("renders business info company type", () => {
    render(ProfilePage, { props: businessProps });
    expect(screen.getByText("Tech")).toBeInTheDocument();
  });
});
