<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  type ProjectDisplay = {
    id: string;
    title: string;
    description: string;
    skills: string[];
    start_time: Date;
    end_time: Date;
    created_by_id: string;
    created_by_name: string;
    status: string;
    max_member: number;
    current_member: number;
    created_at: string;
  };

  const token = data.token;

  let projectsDisplay: ProjectDisplay[] = $state(data.projects);
  let errorLoadingProjects: string | null = $state(data.error);

  let projectsDisplayCopy = projectsDisplay;

  let filterState: string = $state("all");

  // Modal states
  let showEditModal: boolean = $state(false);
  let showDeleteModal: boolean = $state(false);
  let currentProject: ProjectDisplay | null = $state(null);

  // Thêm state cho modal sinh viên đang apply
  let showApplicantsModal: boolean = $state(false);
  let currentApplicants: any[] = $state([]);

  // Mock data cho sinh viên đang apply
  const mockApplicants = [
    {
      id: "user123",
      name: "John Doe",
      major: "Computer Science",
      skills: ["JavaScript", "React"],
      avatar: "https://avatars.githubusercontent.com/u/123456?v=4",
      appliedDate: "2025-04-15",
    },
    {
      id: "user456",
      name: "Jane Smith",
      major: "Data Science & Analytics",
      skills: ["Python", "Data Science"],
      avatar: "https://avatars.githubusercontent.com/u/234567?v=4",
      appliedDate: "2025-04-18",
    },
    {
      id: "user789",
      name: "Robert Chen",
      major: "Software Engineering",
      skills: ["Java", "Spring Boot"],
      avatar: "https://avatars.githubusercontent.com/u/345678?v=4",
      appliedDate: "2025-04-20",
    },
  ];

  // Thêm state cho modal xác nhận xóa
  let showRemoveStudentModal: boolean = $state(false);
  let studentToRemove: any = $state(null);

  // Thêm state cho modal xác nhận rời dự án
  let showLeaveProjectModal: boolean = $state(false);
  let projectToLeave: ProjectDisplay | null = $state(null);

  // Hàm mở modal xác nhận rời dự án
  function openLeaveProjectModal(project: ProjectDisplay) {
    projectToLeave = project;
    showLeaveProjectModal = true;
  }

  // Hàm đóng modal xác nhận rời dự án
  function closeLeaveProjectModal() {
    showLeaveProjectModal = false;
    projectToLeave = null;
  }

  // Hàm xử lý rời dự án
  const leaveProject = async () => {
    if (!projectToLeave) return;

    try {
      // TODO: Thay thế bằng API call thực tế
      // const response = await fetch(`/api/projects/${projectToLeave.id}/leave`, {
      //   method: 'POST',
      //   headers: {
      //     'Authorization': `Bearer ${token}`
      //   }
      // });

      // if (!response.ok) {
      //   const errorData = await response.json();
      //   throw new Error(errorData.error || 'Failed to leave project');
      // }

      // Mock xử lý: Giả lập xóa dự án khỏi danh sách hiển thị hoặc thay đổi trạng thái
      projectsDisplay = projectsDisplay.filter(
        (p) => p.id !== projectToLeave?.id
      );
      projectsDisplayCopy = projectsDisplay;

      // Đóng modal xác nhận
      closeLeaveProjectModal();

      // Hiển thị thông báo thành công
      alert(`You have successfully left the project "${projectToLeave.title}"`);
    } catch (error) {
      console.error("Error leaving project:", error);
      alert("Failed to leave project: " + error);
    }
  };

  // Hàm mở modal xác nhận xóa
  function openRemoveStudentModal(e: Event, student: any, projectId: string) {
    e.preventDefault(); // Ngăn chặn sự kiện nổi bọt
    e.stopPropagation(); // Ngăn chặn click event truyền lên thẻ cha
    studentToRemove = { ...student, projectId };
    showRemoveStudentModal = true;
  }

  function closeRemoveStudentModal() {
    showRemoveStudentModal = false;
    studentToRemove = null;
  }

  // Hàm xử lý xóa sinh viên khỏi dự án
  const removeStudentFromProject = async () => {
    if (!studentToRemove) return;

    try {
      // TODO: Thay thế bằng API call thực tế
      // const response = await fetch(`/api/projects/${studentToRemove.projectId}/students/${studentToRemove.id}`, {
      //   method: 'DELETE',
      //   headers: {
      //     'Authorization': `Bearer ${token}`
      //   }
      // });

      // if (!response.ok) {
      //   const errorData = await response.json();
      //   throw new Error(errorData.error || 'Failed to remove student');
      // }

      // Xóa sinh viên khỏi danh sách hiện tại
      currentApplicants = currentApplicants.filter(
        (a) => a.id !== studentToRemove.id
      );

      // Đóng modal xác nhận
      closeRemoveStudentModal();

      // Hiển thị thông báo thành công
      alert(`Removed ${studentToRemove.name} from the project successfully`);
    } catch (error) {
      console.error("Error removing student:", error);
      alert("Failed to remove student: " + error);
    }
  };

  // Hàm mở modal các sinh viên đang apply
  function openApplicantsModal(projectId: string) {
    // TODO: Thực tế cần fetch danh sách sinh viên đang apply từ API
    // const fetchApplicants = async (projectId) => {
    //   try {
    //     const response = await fetch(`/api/projects/${projectId}/applicants`, {
    //       headers: {
    //         Authorization: `Bearer ${token}`
    //       }
    //     });
    //     if (!response.ok) throw new Error('Failed to fetch applicants');
    //     const applicants = await response.json();
    //     return applicants;
    //   } catch (error) {
    //     console.error('Error fetching applicants:', error);
    //     return [];
    //   }
    // };

    // Hiện tại dùng mock data
    currentApplicants = mockApplicants;
    showApplicantsModal = true;
  }

  function closeApplicantsModal() {
    showApplicantsModal = false;
    currentApplicants = [];
  }

  // Hàm format ngày apply
  function formatApplyDate(dateString: string): string {
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  }

  // Modal functions
  function openEditModal(project: ProjectDisplay) {
    currentProject = { ...project };
    showEditModal = true;
  }

  function openDeleteModal(project: ProjectDisplay) {
    currentProject = project;
    showDeleteModal = true;
  }

  function closeEditModal() {
    showEditModal = false;
    currentProject = null;
  }

  function closeDeleteModal() {
    showDeleteModal = false;
    currentProject = null;
  }

  // Edit form state
  let editTitle: string = $state("");
  let editDescription: string = $state("");
  let editSkills: string[] = $state([]);
  let editStartDate: string = $state("");
  let editEndDate: string = $state("");
  let editMaxMember: number = $state(0);
  let editStatus: string = $state("");

  $effect(() => {
    if (currentProject && showEditModal) {
      editTitle = currentProject.title;
      editDescription = currentProject.description;
      editSkills = currentProject.skills;
      editStartDate = new Date(currentProject.start_time)
        .toISOString()
        .split("T")[0];
      editEndDate = new Date(currentProject.end_time)
        .toISOString()
        .split("T")[0];
      editMaxMember = currentProject.max_member;
      editStatus = currentProject.status;
    }
  });

  // Thêm một computed state để lọc projects dựa trên filterState
  $effect(() => {
    if (filterState === "all") {
      projectsDisplay = projectsDisplayCopy;
    } else if (filterState === "open") {
      projectsDisplay = projectsDisplayCopy.filter(
        (project) => project.status.toLowerCase() === "open"
      );
    } else if (filterState === "close") {
      projectsDisplay = projectsDisplayCopy.filter(
        (project) => project.status.toLowerCase() === "close"
      );
    }
  });

  // Format date function
  function formatDate(dateString: string): string {
    if (!dateString) return "N/A";

    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "Invalid date";

    // Format: "Mar 24, 2025"
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  }

  // Updated to use our server routes
  const deleteProject = async () => {
    try {
      if (!currentProject) return;

      const response = await fetch(`/api/projects/${currentProject.id}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to delete project");
      }

      // Remove the deleted project from the display list
      projectsDisplay = projectsDisplay.filter(
        (p) => p.id !== currentProject?.id
      );
      projectsDisplayCopy = projectsDisplay;

      // Close the modal after successful deletion
      closeDeleteModal();
    } catch (error) {
      console.error("Error deleting project:", error);
      alert("Failed to delete project: " + error);
    }
  };

  // Updated to use our server routes
  const updateProject = async () => {
    try {
      if (!currentProject) return;

      const response = await fetch(`/api/projects/${currentProject.id}`, {
        method: "PUT",
        body: JSON.stringify({
          title: editTitle,
          description: editDescription,
          skills: editSkills,
          start_time: new Date(editStartDate),
          end_time: new Date(editEndDate),
          max_member: editMaxMember,
          status: editStatus,
        }),
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to update project");
      }

      // Update the local state with the edited project
      const updatedProject = await response.json();
      projectsDisplay = projectsDisplay.map((p) =>
        p.id === currentProject?.id ? updatedProject : p
      );

      // Close the modal after successful update
      closeEditModal();
    } catch (error) {
      console.error("Error updating project:", error);
      alert("Failed to update project: " + error);
    }
  };
</script>

<header class="flex justify-between items-center mb-4 ml-64 pr-4 pl-4 pt-4">
  <button class="btn-secondary">
    <a href="/project/application">
      <div class="flex items-center">
        <div
          class="flex items-center justify-center w-5 h-5 bg-red-500 text-white text-xs font-medium rounded-full mr-2"
        >
          <!-- Cân nhắc thêm số lượng application còn pending vào đây nếu còn thời
          gian -->
        </div>
        Applications
      </div>
    </a>
  </button>

  {#if data.role === "business"}
    <button class="btn">
      <a href="/project/create">
        <div class="flex items-center">
          <svg
            class="w-4 h-4 mr-1"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 6v6m0 0v6m0-6h6m-6 0H6"
            ></path>
          </svg>
          Create Project
        </div>
      </a>
    </button>
  {/if}
</header>

<main class="flex-1 pr-4 pl-4 ml-64 pt-2">
  <div class="flex space-x-4">
    <!-- Phần danh sách dự án active -->
    <div class="w-3/5 space-y-3">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">Your Active Projects</h3>
        <div class="flex space-x-2">
          <button
            class={filterState === "all"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "all")}
          >
            All
          </button>
          <button
            class={filterState === "open"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "open")}
          >
            Open
          </button>
          <button
            class={filterState === "close"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "close")}
          >
            Close
          </button>
        </div>
      </div>
      {#if projectsDisplay.length !== 0 && !errorLoadingProjects}
        {#each projectsDisplay as ProjectDisplay}
          <!-- Project Card demo -->
          <a
            href={"/project/" + ProjectDisplay.id}
            class="card p-3 block relative group"
          >
            <div class="flex justify-between items-center">
              <div>
                <h4 class="text-base font-medium">{ProjectDisplay.title}</h4>
                <p class="text-sm text-gray-500">
                  Skills: {ProjectDisplay.skills}
                </p>
                <p class="text-xs text-gray-400 mt-1">
                  Progress: <span class="accent-color">60%</span> | Start: {formatDate(
                    ProjectDisplay.start_time.toString()
                  )}
                  | End: {formatDate(ProjectDisplay.end_time.toString())}
                </p>
              </div>

              <div class="flex space-x-2 items-center">
                <!-- Hiển thị số lượng thành viên mới - thiết kế đẹp hơn -->
                <div class="flex items-center bg-gray-100 px-2 py-1 rounded-lg">
                  <div class="flex -space-x-2 mr-2">
                    <!-- Hiển thị avatar đại diện cho thành viên (có thể là placeholder) -->
                    <div
                      class="w-6 h-6 rounded-full bg-purple-500 border-2 border-white flex items-center justify-center"
                    >
                      <span class="text-[10px] font-bold text-white"
                        >{ProjectDisplay.current_member > 0
                          ? ProjectDisplay.current_member
                          : ""}</span
                      >
                    </div>
                    {#if ProjectDisplay.current_member >= 2}
                      <div
                        class="w-6 h-6 rounded-full bg-indigo-500 border-2 border-white flex items-center justify-center"
                      ></div>
                    {/if}
                    {#if ProjectDisplay.current_member >= 3}
                      <div
                        class="w-6 h-6 rounded-full bg-blue-500 border-2 border-white flex items-center justify-center"
                      ></div>
                    {/if}
                  </div>

                  <div class="flex items-center">
                    <span class="text-xs font-medium">
                      {ProjectDisplay.current_member}/{ProjectDisplay.max_member}
                    </span>
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 24 24"
                      fill="currentColor"
                      class="w-4 h-4 ml-1 text-gray-500"
                    >
                      <path
                        d="M4.5 6.375a4.125 4.125 0 118.25 0 4.125 4.125 0 01-8.25 0zM14.25 8.625a3.375 3.375 0 116.75 0 3.375 3.375 0 01-6.75 0zM1.5 19.125a7.125 7.125 0 0114.25 0v.003l-.001.119a.75.75 0 01-.363.63 13.067 13.067 0 01-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 01-.364-.63l-.001-.122zM17.25 19.128l-.001.144a2.25 2.25 0 01-.233.96 10.088 10.088 0 005.06-1.01.75.75 0 00.42-.643 4.875 4.875 0 00-6.957-4.611 8.586 8.586 0 011.71 5.157v.003z"
                      />
                    </svg>
                  </div>
                </div>

                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded font-bold"
                  >{ProjectDisplay.status.toUpperCase()}</span
                >

                <svg
                  class="w-5 h-5 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                  ></path>
                </svg>
              </div>
            </div>
          </a>
          <!-- Project action buttons - Only show for business users and only on hover -->
          {#if data.role === "business"}
            <div
              class="group-hover:opacity-100 transition-opacity flex space-x-2"
            >
              <button
                class="p-1.5 bg-blue-100 text-blue-600 rounded hover:bg-blue-200"
                onclick={() => openEditModal(ProjectDisplay)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                  />
                </svg>
              </button>

              <!-- Nút mới - Quản lý ứng viên -->
              <button
                class="p-1.5 bg-purple-100 text-purple-600 rounded hover:bg-purple-200"
                onclick={(e) => {
                  e.preventDefault(); // Ngăn chặn chuyển hướng từ thẻ a
                  openApplicantsModal(ProjectDisplay.id);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                  />
                </svg>
              </button>

              <button
                class="p-1.5 bg-red-100 text-red-600 rounded hover:bg-red-200"
                onclick={() => {
                  openDeleteModal(ProjectDisplay);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                  />
                </svg>
              </button>
            </div>
          {:else if data.role === "student"}
            <div
              class="group-hover:opacity-100 transition-opacity flex space-x-2"
            >
              <button
                class="p-1.5 bg-purple-100 text-purple-600 rounded hover:bg-purple-200"
                onclick={(e) => {
                  e.preventDefault(); // Ngăn chặn chuyển hướng từ thẻ a
                  openApplicantsModal(ProjectDisplay.id);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                  />
                </svg>
              </button>
              <button
                class="p-1.5 bg-red-100 text-red-600 rounded hover:bg-red-200"
                onclick={(e) => {
                  e.preventDefault();
                  openLeaveProjectModal(ProjectDisplay);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                  />
                </svg>
              </button>
            </div>
          {/if}
        {/each}
      {:else}
        <div class="card p-4 text-center">
          <p class="text-gray-500">No projects found.</p>
          <!-- <button class="btn mt-2">Create Your First Project</button> -->
        </div>
      {/if}
    </div>

    {#if data.role === "business"}
      <!-- Right sidebar with analytics and talent pool -->
      <div class="w-2/5 space-y-4">
        <!-- Business Analytics -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Project Analytics</h3>
          <div class="space-y-4">
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Active Projects</span>
                <span class="text-sm font-medium">4</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-[#6b48ff]" style="width: 80%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Talent Allocation</span>
                <span class="text-sm font-medium">12/15</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-[#ff6f61]" style="width: 75%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Tasks Completion</span>
                <span class="text-sm font-medium">92%</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-green-500" style="width: 92%"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Active Talent Pool -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Talent Pool</h3>
          <div class="space-y-2">
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/234567?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Sarah Johnson</p>
                  <p class="text-xs text-gray-500">React, Node.js</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/345678?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Michael Chen</p>
                  <p class="text-xs text-gray-500">Python, MongoDB</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/456789?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Alex Rivera</p>
                  <p class="text-xs text-gray-500">UI/UX, Vue.js</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div class="flex justify-center mt-2">
              <button class="text-sm text-[#6b48ff] hover:underline">
                View All Talent
              </button>
            </div>
          </div>
        </div>
      </div>
    {/if}
  </div>

  {#if showApplicantsModal}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div
        class="bg-white rounded-lg p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto"
      >
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-xl font-semibold">Project Applicants</h2>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeApplicantsModal}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>

        {#if currentApplicants.length > 0}
          <div class="space-y-4">
            {#each currentApplicants as applicant}
              <div
                class=" rounded-lg p-4 flex items-center justify-between bg-gray-100 hover:bg-gray-100 transition-colors"
              >
                <div class="flex items-center space-x-4">
                  <img
                    src={applicant.avatar}
                    alt={applicant.name}
                    class="w-12 h-12 rounded-full object-cover border-2 border-purple-200"
                  />
                  <div>
                    <h3 class="font-medium">{applicant.name}</h3>
                    <!-- Thêm hiển thị chuyên ngành -->
                    <p class="text-xs text-gray-600 mt-0.5">
                      {applicant.major}
                    </p>
                    <p class="text-sm text-gray-500">
                      Skills: {applicant.skills.join(", ")}
                    </p>
                    <p class="text-xs text-gray-400">
                      Applied: {formatApplyDate(applicant.appliedDate)}
                    </p>
                  </div>
                </div>
                <div class="flex space-x-2">
                  <!-- Nút xem hồ sơ sinh viên -->
                  <a
                    href={`/profile/${applicant.id}`}
                    class="px-3 py-1.5 bg-[#6b48ff] text-white text-sm rounded hover:bg-[#5a3dd3]"
                  >
                    View Profile
                  </a>

                  <!-- Thêm nút xóa sinh viên -->
                  <button
                    class="px-3 py-1.5 bg-red-100 text-red-600 text-sm rounded hover:bg-red-200"
                    onclick={(e) =>
                      openRemoveStudentModal(
                        e,
                        applicant,
                        currentProject?.id || ""
                      )}
                  >
                    Remove
                  </button>
                </div>
              </div>
            {/each}
          </div>
        {:else}
          <div class="text-center py-8">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-12 w-12 mx-auto text-gray-400"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"
              />
            </svg>
            <p class="mt-4 text-gray-600">No applicants yet</p>
          </div>
        {/if}
      </div>
    </div>
  {/if}
  <!-- Edit Project Modal -->
  {#if showEditModal && currentProject}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-2xl max-h-[90vh]">
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-xl font-semibold">Edit Project</h2>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeEditModal}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>

        <form class="space-y-4">
          <div>
            <label for="title" class="block text-sm font-medium text-gray-700"
              >Project Title</label
            >
            <input
              type="text"
              id="title"
              bind:value={editTitle}
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
          </div>

          <div>
            <label
              for="description"
              class="block text-sm font-medium text-gray-700">Description</label
            >
            <textarea
              id="description"
              bind:value={editDescription}
              rows="4"
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            ></textarea>
          </div>

          <div>
            <label for="skills" class="block text-sm font-medium text-gray-700"
              >Skills (comma separated)</label
            >
            <input
              type="text"
              id="skills"
              bind:value={editSkills}
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label
                for="startDate"
                class="block text-sm font-medium text-gray-700"
                >Start Date</label
              >
              <input
                type="date"
                id="startDate"
                bind:value={editStartDate}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>

            <div>
              <label
                for="endDate"
                class="block text-sm font-medium text-gray-700">End Date</label
              >
              <input
                type="date"
                id="endDate"
                bind:value={editEndDate}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label
                for="maxMember"
                class="block text-sm font-medium text-gray-700"
                >Max Members</label
              >
              <input
                type="number"
                id="maxMember"
                bind:value={editMaxMember}
                min={currentProject.current_member}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>

            <div>
              <label
                for="status"
                class="block text-sm font-medium text-gray-700">Status</label
              >
              <select
                id="status"
                bind:value={editStatus}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              >
                <option value="open">Open</option>
                <option value="close">Closed</option>
              </select>
            </div>
          </div>

          <div class="flex justify-end space-x-2 pt-4">
            <button
              type="button"
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeEditModal}
            >
              Cancel
            </button>
            <button
              type="button"
              class="px-4 py-2 bg-[#6b48ff] text-white rounded-md hover:bg-[#5a3dd3]"
              onclick={updateProject}
            >
              Save Changes
            </button>
          </div>
        </form>
      </div>
    </div>
  {/if}

  <!-- Delete Project Modal -->
  {#if showDeleteModal && currentProject}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <div class="text-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-12 w-12 mx-auto text-red-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
            />
          </svg>

          <h2 class="text-xl font-semibold mt-4">Delete Project</h2>

          <p class="mt-2 text-gray-600">
            Are you sure you want to delete "{currentProject.title}"? This
            action cannot be undone.
          </p>

          <div class="flex justify-center space-x-3 mt-6">
            <button
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeDeleteModal}
            >
              Cancel
            </button>
            <button
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
              onclick={deleteProject}
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}

  <!-- Thêm modal xác nhận xóa sinh viên -->
  {#if showRemoveStudentModal && studentToRemove}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <div class="text-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-12 w-12 mx-auto text-red-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
            />
          </svg>

          <h2 class="text-xl font-semibold mt-4">Remove Team Member</h2>

          <p class="mt-2 text-gray-600">
            Are you sure you want to remove <span class="font-medium"
              ><strong>{studentToRemove.name}</strong></span
            >
            from this project? This action <strong>cannot be undone</strong>.
          </p>

          <div class="flex justify-center space-x-3 mt-6">
            <button
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeRemoveStudentModal}
            >
              Cancel
            </button>
            <button
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
              onclick={removeStudentFromProject}
            >
              Remove
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}

  <!-- Leave Project Modal -->
  {#if showLeaveProjectModal && projectToLeave}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <div class="text-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-12 w-12 mx-auto text-red-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
            />
          </svg>

          <h2 class="text-xl font-semibold mt-4">Leave Project</h2>

          <p class="mt-2 text-gray-600">
            Are you sure you want to leave <span class="font-medium"
              ><strong>{projectToLeave.title}</strong></span
            >? This action <strong>cannot be undone</strong>.
          </p>

          <div class="flex justify-center space-x-3 mt-6">
            <button
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeLeaveProjectModal}
            >
              Cancel
            </button>
            <button
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
              onclick={leaveProject}
            >
              Leave Project
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}
</main>

<style>
  [contenteditable]:empty:before {
    content: attr(placeholder);
    color: #9ca3af;
    display: block;
  }
</style>
