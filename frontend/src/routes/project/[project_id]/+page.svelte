<script lang="ts">
  import type { PageData } from "./$types";
  import { onMount, onDestroy } from "svelte";
  import type { Socket } from "socket.io-client";
  import { browser } from "$app/environment";

  let { data }: { data: PageData } = $props();
  const project = data.project;
  const userId = data.id;
  const token = data.token;

  let socket: Socket;
  let tasks = $state<
    { id: string; title: string; status: string; assigned_to?: string }[]
  >([]);
  let activities = $state<
    {
      id: number;
      user: string;
      action: string;
      item: string;
      targetStatus?: string;
      timestamp: string;
    }[]
  >([]);
  let sortables: any[] = [];

  // Khởi tạo SortableJS cho các cột trong Kanban
  async function initSortable() {
    if (!browser) return;

    try {
      // Lazy load Sortable
      const SortableModule = await import("sortablejs");
      const Sortable = SortableModule.default;

      // Xóa các instance cũ để tránh duplicate
      sortables.forEach((sortable) => sortable.destroy());
      sortables = [];

      const columns = document.querySelectorAll(".kanban-column");
      columns.forEach((column) => {
        const status = column.getAttribute("data-status");
        const sortable = new Sortable(column as HTMLElement, {
          group: "tasks",
          animation: 150,
          ghostClass: "task-ghost",
          chosenClass: "task-chosen",
          dragClass: "task-drag",
          filter: ".ignore-elements", // Phần tử không kéo được
          onEnd: function (evt) {
            const taskId = evt.item.getAttribute("data-id");
            const fromStatus = evt.from.getAttribute("data-status");
            const toStatus = evt.to.getAttribute("data-status");

            // Nếu di chuyển sang cột khác
            if (fromStatus !== toStatus) {
              if (taskId && fromStatus && toStatus)
                handleTaskMove(taskId, fromStatus, toStatus);
            }
          },
        });
        sortables.push(sortable);
      });
    } catch (error) {
      console.error("Error loading Sortable:", error);
    }
  }

  function handleTaskMove(
    taskId: string,
    fromStatus: string,
    toStatus: string
  ) {
    if (socket) {
      console.log(`Moving task: ${taskId} from ${fromStatus} to ${toStatus}`);

      // Cập nhật local trước - Optimistic UI
      updateTaskPosition(taskId, toStatus);

      // Gửi yêu cầu cập nhật đến server
      socket.emit("message", {
        type: "update",
        content: {
          taskId: taskId,
          status: toStatus,
        },
      });
    }
  }

  function updateTaskPosition(taskId: string, newStatus: string) {
    // Cập nhật vị trí task trong UI
    tasks = tasks.map((task) => {
      if (task.id === taskId) {
        return { ...task, status: newStatus };
      }
      return task;
    });
  }

  // Lọc task theo trạng thái
  function getTasksByStatus(status: string) {
    return tasks.filter((task) => task.status === status);
  }

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

  // Calculate project progress based on start and end dates
  function calculateProgress(start: string, end: string): number {
    const startDate = new Date(start);
    const endDate = new Date(end);
    const currentDate = new Date();

    // If dates are invalid, return 0
    if (isNaN(startDate.getTime()) || isNaN(endDate.getTime())) {
      return 0;
    }

    // If project hasn't started yet
    if (currentDate < startDate) {
      return 0;
    }

    // If project is completed
    if (currentDate > endDate) {
      return 100;
    }

    // Calculate progress percentage
    const totalDuration = endDate.getTime() - startDate.getTime();
    const elapsedDuration = currentDate.getTime() - startDate.getTime();
    const progress = Math.floor((elapsedDuration / totalDuration) * 100);

    return progress;
  }

  // Calculate progress percentage
  const progressPercentage = calculateProgress(
    project.start_time,
    project.end_time
  );

  // Task creation modal
  let showTaskModal = $state(false);
  let newTasks = $state([
    { name: "", description: "", note: "", assignedTo: "" },
  ]);

  function addAnotherTask() {
    newTasks.push({ name: "", description: "", note: "", assignedTo: "" });
  }

  function removeTask(index: number) {
    newTasks = newTasks.filter((_, i) => i !== index);
  }

  function openTaskModal() {
    showTaskModal = true;
  }

  function closeTaskModal() {
    showTaskModal = false;
    // Reset tasks form
    newTasks = [{ name: "", description: "", note: "", assignedTo: "" }];
  }

  function createTasksWithConfirm() {
    // Filter out empty tasks
    const validTasks = newTasks.filter((task) => task.name.trim() !== "");

    if (validTasks.length === 0) {
      alert("Please add at least one task with a name");
      return;
    }

    // Tạo từng task và gửi lên server
    validTasks.forEach((task) => {
      createTask(task);
    });

    closeTaskModal();
  }

  async function fetchTeamMembers() {
    try {
      const response = await fetch(`/api/projects/${project.id}/students`);

      if (!response.ok) {
        throw new Error("Failed to fetch team members");
      }

      const data = await response.json();
      console.log("Team members from API:", data);

      // Chuyển đổi dữ liệu từ API sang định dạng hiển thị
      teamMembers = data.map((student: any) => ({
        id: student.id,
        name: student.name || student.username,
        role: student.role || "Member",
        email: student.email || "",
      }));
    } catch (error) {
      console.error("Error fetching team members:", error);
    }
  }
  // Gửi yêu cầu tạo task mới đến server
  function createTask(taskInput: {
    name: string;
    description: string;
    note?: string;
    assignedTo?: string;
  }) {
    if (socket) {
      const message = {
        type: "create",
        content: {
          title: taskInput.name,
          description: taskInput.description,
          note: taskInput.note || "",
          assigned_to: taskInput.assignedTo || "",
        },
      };
      console.log("Sending create task message:", message);

      try {
        socket.emit("message", message);
        console.log("Message sent successfully");
      } catch (error) {
        console.error("Error sending message:", error);
      }
    } else {
      console.error("Socket not connected");
    }
  }

  // Activities modal
  let showActivitiesModal = $state(false);

  // Mock data for activities
  const mockActivities = $state([
    {
      id: 1,
      user: "Sarah",
      action: "moved",
      item: "Test Features",
      targetStatus: "Review",
      timestamp: "2025-04-24T14:32:00",
    },
    {
      id: 2,
      user: "John",
      action: "moved",
      item: "Code Backend",
      targetStatus: "In Progress",
      timestamp: "2025-04-24T13:45:00",
    },
    {
      id: 3,
      user: "Alex",
      action: "completed",
      item: "Login Screen UI",
      targetStatus: "Done",
      timestamp: "2025-04-24T11:20:00",
    },
    {
      id: 4,
      user: "Emma",
      action: "created",
      item: "New Feature Request",
      targetStatus: "To-do",
      timestamp: "2025-04-23T16:15:00",
    },
    {
      id: 5,
      user: "Michael",
      action: "commented on",
      item: "API Documentation",
      targetStatus: "",
      timestamp: "2025-04-23T14:32:00",
    },
    {
      id: 6,
      user: "Tran Le Minh Nhat",
      action: "assigned",
      item: "Wireframes",
      targetStatus: "To-do",
      timestamp: "2025-04-22T09:30:00",
    },
    {
      id: 7,
      user: "Nguyen My Thong",
      action: "updated",
      item: "Authentication Logic",
      targetStatus: "",
      timestamp: "2025-04-21T17:25:00",
    },
    {
      id: 8,
      user: "Tran Tuan Kiet",
      action: "completed",
      item: "Database Schema",
      targetStatus: "Done",
      timestamp: "2025-04-21T15:10:00",
    },
    {
      id: 9,
      user: "Tran Le Minh Nhat",
      action: "created",
      item: "Project Repository",
      targetStatus: "",
      timestamp: "2025-04-20T11:45:00",
    },
    {
      id: 10,
      user: "Sarah",
      action: "added",
      item: "New Team Member",
      targetStatus: "",
      timestamp: "2025-04-19T10:30:00",
    },
  ]);

  // Team Members modal
  let showTeamMembersModal = $state(false);

  function openTeamMembersModal() {
    showTeamMembersModal = true;
  }

  function closeTeamMembersModal() {
    showTeamMembersModal = false;
  }

  // Extend the team members data with more info
  let teamMembers = $state<
    {
      id: number | string;
      name: string;
      role: string;
      email: string;
    }[]
  >([]);

  // Function to format the date for grouping
  function formatActivityDate(dateString: string): string {
    const date = new Date(dateString);

    // Get today and yesterday dates for comparison
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    const yesterday = new Date(today);
    yesterday.setDate(yesterday.getDate() - 1);

    // Check if the date is today, yesterday, or earlier
    if (date >= today) {
      return "Today";
    } else if (date >= yesterday) {
      return "Yesterday";
    } else {
      return date.toLocaleDateString("en-US", {
        month: "short",
        day: "numeric",
        year: "numeric",
      });
    }
  }

  // Function to format the time
  function formatActivityTime(dateString: string): string {
    const date = new Date(dateString);
    return date.toLocaleTimeString("en-US", {
      hour: "2-digit",
      minute: "2-digit",
    });
  }

  // Function to get the grouped activities
  function getGroupedActivities() {
    // Sử dụng activities từ server nếu có, nếu không dùng mockActivities
    const activitiesData = activities.length > 0 ? activities : mockActivities;
    const groups: Record<string, any[]> = {};

    activitiesData.forEach((activity) => {
      const timestamp = activity.timestamp;
      const dateGroup = formatActivityDate(timestamp);

      if (!groups[dateGroup]) {
        groups[dateGroup] = [];
      }
      groups[dateGroup].push(activity);
    });

    // Convert the groups object to an array of { date, activities } objects
    return Object.entries(groups).map(([date, activities]) => {
      return { date, activities };
    });
  }

  // // Hàm xử lý activity từ server
  // function formatActivityFromServer(activity) {
  //   // Chuyển đổi cấu trúc activity từ server sang định dạng hiển thị
  //   // Dạng activity từ server: {id, type, done_by, project_id, title, from, to, created_at}

  //   let action = "";
  //   switch (activity.type) {
  //     case "create":
  //       action = "created";
  //       break;
  //     case "update":
  //       action = "updated";
  //       break;
  //     case "delete":
  //       action = "deleted";
  //       break;
  //     case "move":
  //       action = "moved";
  //       break;
  //     default:
  //       action = activity.type;
  //   }

  //   return {
  //     id: activity.id,
  //     user: activity.done_by_name || activity.done_by,
  //     action: action,
  //     item: activity.title,
  //     targetStatus: activity.to,
  //     timestamp: activity.created_at,
  //   };
  // }

  function openActivitiesModal() {
    showActivitiesModal = true;
  }

  function closeActivitiesModal() {
    showActivitiesModal = false;
  }

  // Task edit modal
  let showEditTaskModal = $state(false);
  let currentEditingTask = $state<{
    id: string;
    title: string;
    description: string;
    note?: string;
    status: string;
    assigned_to?: string;
  } | null>(null);

  function openEditTaskModal(task: any) {
    currentEditingTask = { ...task };
    showEditTaskModal = true;
  }

  function closeEditTaskModal() {
    showEditTaskModal = false;
    currentEditingTask = null;
  }

  function updateTask() {
    if (socket && currentEditingTask) {
      socket.emit("message", {
        type: "update",
        content: {
          taskId: currentEditingTask.id,
          title: currentEditingTask.title,
          description: currentEditingTask.description,
          note: currentEditingTask.note || "",
          status: currentEditingTask.status,
          assigned_to: currentEditingTask.assigned_to || "",
        },
      });
      closeEditTaskModal();
    }
  }

  function deleteTask() {
    if (socket && currentEditingTask) {
      if (confirm("Are you sure you want to delete this task?")) {
        socket.emit("message", {
          type: "delete",
          content: currentEditingTask.id,
        });
        closeEditTaskModal();
      }
    }
  }

  onMount(async () => {
    if (browser) {
      try {
        await fetchTeamMembers();
        // Sửa URL để khớp với API của backend
        const wsUrl = `ws://${window.location.hostname}:8080/ws/task/${project.id}/${data.id}`;
        console.log("Connecting with WebSocket to:", wsUrl);

        // Tạo một webSocket
        const ws = new WebSocket(wsUrl);

        ws.onopen = () => {
          console.log("WebSocket connection established");
        };

        ws.onmessage = (event) => {
          try {
            const data = JSON.parse(event.data);
            console.log("Received via WebSocket:", data);

            if (data.tasks) {
              tasks = data.tasks;
              setTimeout(initSortable, 100);
            }

            if (data.activities) {
              activities = data.activities;
            }
          } catch (e) {
            console.error("Error parsing WebSocket data:", e);
          }
        };

        ws.onerror = (err) => {
          console.error("WebSocket error:", err);
        };

        // Sửa hàm emit để gửi message đúng định dạng
        socket = {
          emit: (eventName: string, data: any) => {
            if (ws.readyState === WebSocket.OPEN) {
              // Gửi tin nhắn trực tiếp theo định dạng backend yêu cầu
              // Backend mong đợi { type: "...", content: {...} }
              ws.send(
                JSON.stringify({
                  type: data.type, // lấy type từ data object
                  content: data.content, // lấy content từ data object
                })
              );
            } else {
              console.error("WebSocket is not connected");
            }
          },
          disconnect: () => {
            ws.close();
          },
        } as any;

        // Khởi tạo sortable
        setTimeout(initSortable, 100);
      } catch (error) {
        console.error("Error setting up WebSocket:", error);
      }
    }
  });

  onDestroy(() => {
    if (socket) {
      socket.disconnect();
    }

    // Dọn dẹp sortable instances
    sortables.forEach((sortable) => sortable.destroy());
  });
</script>

<header class="flex justify-between items-center ml-64 pr-4 pl-4 pt-4">
  <div class="flex items-center mb-4">
    <a href="/project">
      <button class="text-gray-500 hover:text-gray-700 mr-3">
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M10 19l-7-7m0 0l7-7m-7 7h18"
          ></path>
        </svg>
      </button>
    </a>
    <div>
      <h2 class="text-xl font-semibold">{project.title}</h2>
      <p class="text-sm text-gray-600">Manage your project</p>
    </div>
  </div>
</header>

<main class="flex-1 pr-4 pl-4 ml-64">
  <div class="space-y-4">
    <!-- Top section: Timeline, Progression, and Recent Activities -->
    <div class="flex space-x-4">
      <!-- NEW: Project Info & Team Members Card -->
      <div class="w-72">
        <!-- NEW: Team Members Card -->
        <div class="card p-4">
          <div class="flex justify-between items-center mb-3">
            <h3 class="text-base font-semibold">Team Members</h3>
            <button
              class="text-xs text-blue-600 hover:underline"
              onclick={openTeamMembersModal}
            >
              View all
            </button>
          </div>

          <div class="space-y-3">
            {#each teamMembers.slice(0, 3) as member}
              <div class="flex items-center space-x-2">
                <div
                  class="w-8 h-8 rounded-full bg-gray-200 flex items-center justify-center text-sm font-medium text-gray-600"
                >
                  {member.name
                    .split(" ")
                    .map((name) => name[0])
                    .join("")}
                </div>
                <div>
                  <p class="text-sm font-medium">{member.name}</p>
                  <p class="text-xs text-gray-500">Developer</p>
                </div>
              </div>
            {/each}

            {#if teamMembers.length > 3}
              <button
                class="text-xs text-gray-500 hover:text-[#6b48ff]"
                onclick={openTeamMembersModal}
              >
                +{teamMembers.length - 3} more members
              </button>
            {/if}
          </div>
        </div>
      </div>

      <!-- Timeline and Progression (made smaller and compact) -->
      <div class="flex-1 space-y-4">
        <div class="flex space-x-4">
          <!-- Timeline Bar (smaller) -->
          <div class="card p-3 w-1/2">
            <div class="flex justify-between items-center mb-2">
              <h3 class="text-base font-semibold">Timeline</h3>
              <span class="text-sm accent-color">{progressPercentage}%</span>
            </div>
            <div class="flex items-center justify-between">
              <div class="text-xs text-gray-500">
                {formatDate(project.start_time.toString())}
              </div>
              <div class="flex-1 h-2 bg-gray-200 rounded mx-2 relative">
                <div
                  class="absolute h-2 bg-[#6b48ff] rounded"
                  style="width: {progressPercentage}%"
                ></div>
              </div>
              <div class="text-xs text-gray-500">
                {formatDate(project.end_time.toString())}
              </div>
            </div>
          </div>

          <!-- Progression Bar (smaller) -->
          <div class="card p-3 w-1/2">
            <div class="flex justify-between items-center mb-2">
              <h3 class="text-base font-semibold">Tasks Progress</h3>
              <span class="text-sm text-green-600">35%</span>
            </div>
            <div class="flex items-center justify-between">
              <div class="text-xs text-gray-500">Completed: 4</div>
              <div class="flex-1 h-2 bg-gray-200 rounded mx-2 relative">
                <div
                  class="absolute h-2 bg-green-500 rounded"
                  style="width: 35%"
                ></div>
              </div>
              <div class="text-xs text-gray-500">Total: 12</div>
            </div>
          </div>
        </div>

        <!-- Recent Activities (Keep as is) -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-2">Recent Activities</h3>
          <!-- Timeline -->
          <div class="relative">
            <!-- Timeline line -->
            <div class="absolute left-4 top-0 bottom-0 w-0.5 bg-gray-200"></div>

            <!-- Activity items -->
            <div class="space-y-6">
              <!-- Today's activities -->
              <div class="relative pl-8">
                <!-- Timeline dot -->
                <div
                  class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-[#6b48ff] border-2 border-white"
                ></div>
                <span class="text-xs font-semibold text-gray-500 mb-1 block"
                  >18 Mar, 14:32</span
                >
                <p class="text-sm">
                  <span class="font-medium text-[#6b48ff]">Sarah</span>
                  moved <span class="font-medium">Test Features</span> to Review
                </p>
              </div>

              <div class="relative pl-8">
                <!-- Timeline dot -->
                <div
                  class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-[#6b48ff] border-2 border-white"
                ></div>
                <span class="text-xs font-semibold text-gray-500 mb-1 block"
                  >18 Mar, 13:45</span
                >
                <p class="text-sm">
                  <span class="font-medium text-[#6b48ff]">John</span> moved
                  <span class="font-medium">Code Backend</span> to In Progress
                </p>
              </div>
            </div>
          </div>

          <!-- View all button -->
          <button
            class="btn-secondary w-full mt-4 text-sm"
            onclick={openActivitiesModal}
          >
            View All Activities
          </button>
        </div>
      </div>
    </div>

    <!-- Kanban Board (full width, full height) -->
    <div class="card p-3 w-full min-h-[calc(100vh-380px)] mb-4">
      <h3 class="text-base font-semibold mb-2">Tasks</h3>
      <div class="grid grid-cols-4 gap-3 h-full">
        <!-- To-do Column -->
        <div class="flex flex-col">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-2">
              <div class="p-1.5 rounded-md bg-gray-100">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4 text-gray-600"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <circle cx="12" cy="12" r="10"></circle>
                  <line x1="12" y1="8" x2="12" y2="16"></line>
                  <line x1="8" y1="12" x2="16" y2="12"></line>
                </svg>
              </div>
              <h4 class="text-sm font-medium">To-do</h4>
            </div>
            <span
              class="inline-flex items-center justify-center w-5 h-5 text-xs font-medium rounded-full bg-gray-100 text-gray-700"
              >{getTasksByStatus("todo").length}</span
            >
          </div>
          <button
            class="text-xs text-[#6b48ff] mb-2 ignore-elements"
            onclick={openTaskModal}>+ Add tasks</button
          >
          <div
            class="bg-gray-50 rounded-lg p-2 flex-1 overflow-y-auto kanban-column"
            data-status="todo"
          >
            {#each getTasksByStatus("todo") as task (task.id)}
              <div
                class="card p-2 bg-white mb-2 shadow-sm cursor-move"
                data-id={task.id}
              >
                <div class="flex justify-between items-center">
                  <div class="flex-1">
                    <p class="text-sm font-medium">{task.title}</p>
                    <p class="text-xs text-gray-500">
                      Assign to:
                      <a href="#" class="text-[#6b48ff]">{task.assigned_to}</a>
                    </p>
                  </div>
                  <button
                    class="text-gray-400 hover:text-gray-600 ml-2 p-1 rounded-full hover:bg-gray-100 self-center flex-shrink-0 ignore-elements"
                    onclick={() => openEditTaskModal(task)}
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
                        d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            {/each}
          </div>
        </div>

        <!-- In Progress Column -->
        <div class="flex flex-col">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-2">
              <div class="p-1.5 rounded-md bg-yellow-100">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4 text-yellow-600"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path
                    d="M12 2v4M12 18v4M4.93 4.93l2.83 2.83M16.24 16.24l2.83 2.83M2 12h4M18 12h4M4.93 19.07l2.83-2.83M16.24 7.76l2.83-2.83"
                  ></path>
                </svg>
              </div>
              <h4 class="text-sm font-medium">In Progress</h4>
            </div>
            <span
              class="inline-flex items-center justify-center w-5 h-5 text-xs font-medium rounded-full bg-yellow-100 text-yellow-700"
              >{getTasksByStatus("in_progress").length}</span
            >
          </div>
          <div
            class="bg-yellow-50 rounded-lg p-2 flex-1 overflow-y-auto kanban-column"
            data-status="in_progress"
          >
            {#each getTasksByStatus("in_progress") as task (task.id)}
              <div
                class="card p-2 bg-white mb-2 shadow-sm cursor-move"
                data-id={task.id}
              >
                <div class="flex justify-between items-center">
                  <div class="flex-1">
                    <p class="text-sm font-medium">{task.title}</p>
                    <p class="text-xs text-gray-500">
                      Doing by:
                      <a href="#" class="text-[#6b48ff]">{task.assigned_to}</a>
                    </p>
                  </div>
                  <button
                    class="text-gray-400 hover:text-gray-600 ml-2 p-1 rounded-full hover:bg-gray-100 self-center flex-shrink-0 ignore-elements"
                    onclick={() => openEditTaskModal(task)}
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
                        d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            {/each}
          </div>
        </div>

        <!-- Review Column -->
        <div class="flex flex-col">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-2">
              <div class="p-1.5 rounded-md bg-blue-100">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4 text-blue-600"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"></path>
                  <circle cx="12" cy="12" r="3"></circle>
                </svg>
              </div>
              <h4 class="text-sm font-medium">Review</h4>
            </div>
            <span
              class="inline-flex items-center justify-center w-5 h-5 text-xs font-medium rounded-full bg-blue-100 text-blue-700"
              >{getTasksByStatus("review").length}</span
            >
          </div>
          <div
            class="bg-blue-50 rounded-lg p-2 flex-1 overflow-y-auto kanban-column"
            data-status="review"
          >
            {#each getTasksByStatus("review") as task (task.id)}
              <div
                class="card p-2 bg-white mb-2 shadow-sm cursor-move"
                data-id={task.id}
              >
                <div class="flex justify-between items-center">
                  <div class="flex-1">
                    <p class="text-sm font-medium">{task.title}</p>
                    <p class="text-xs text-gray-500">
                      Done by:
                      <a href="#" class="text-[#6b48ff]">{task.assigned_to}</a>
                    </p>
                  </div>
                  <button
                    class="text-gray-400 hover:text-gray-600 ml-2 p-1 rounded-full hover:bg-gray-100 self-center flex-shrink-0 ignore-elements"
                    onclick={() => openEditTaskModal(task)}
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
                        d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            {/each}
          </div>
        </div>

        <!-- Done Column -->
        <div class="flex flex-col">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-2">
              <div class="p-1.5 rounded-md bg-green-100">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4 text-green-600"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path d="M20 6L9 17l-5-5"></path>
                </svg>
              </div>
              <h4 class="text-sm font-medium">Done</h4>
            </div>
            <span
              class="inline-flex items-center justify-center w-5 h-5 text-xs font-medium rounded-full bg-green-100 text-green-700"
              >{getTasksByStatus("done").length}</span
            >
          </div>
          <div
            class="bg-green-50 rounded-lg p-2 flex-1 overflow-y-auto kanban-column"
            data-status="done"
          >
            {#each getTasksByStatus("done") as task (task.id)}
              <div
                class="card p-2 bg-white mb-2 shadow-sm cursor-move"
                data-id={task.id}
              >
                <div class="flex justify-between items-center">
                  <div class="flex-1">
                    <p class="text-sm font-medium">{task.title}</p>
                    <p class="text-xs text-gray-500">
                      Completed by:
                      <a href="#" class="text-[#6b48ff]">{task.assigned_to}</a>
                    </p>
                  </div>
                  <button
                    class="text-gray-400 hover:text-gray-600 ml-2 p-1 rounded-full hover:bg-gray-100 self-center flex-shrink-0 ignore-elements"
                    onclick={() => openEditTaskModal(task)}
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
                        d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            {/each}
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Activities Modal -->
  {#if showActivitiesModal}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center modal"
    >
      <div
        class="bg-white rounded-lg w-full max-w-2xl max-h-[90vh] flex flex-col overflow-hidden shadow-lg"
      >
        <div class="p-4 border-b border-gray-100 bg-white flex-shrink-0">
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-semibold">Project Activities</h3>
            <button
              class="text-gray-500 hover:text-gray-700"
              onclick={closeActivitiesModal}
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
        </div>

        <!-- Content - phần này sẽ cuộn -->
        <div class="p-4 pb-6 overflow-y-auto flex-grow">
          {#each getGroupedActivities() as group}
            <div class="mb-6">
              <h4 class="text-sm font-semibold text-gray-500 mb-3">
                {group.date}
              </h4>

              <!-- Timeline -->
              <div class="relative">
                <!-- Timeline line -->
                <div
                  class="absolute left-4 top-0 bottom-0 w-0.5 bg-gray-200"
                ></div>

                <!-- Activity items -->
                <div class="space-y-4">
                  {#each group.activities as activity}
                    <div class="relative pl-10">
                      <!-- Timeline dot -->
                      <div
                        class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-[#6b48ff] border-2 border-white"
                      ></div>

                      <div class="flex justify-between items-start">
                        <div>
                          <span
                            class="text-xs font-semibold text-gray-500 mb-1 block"
                          >
                            {formatActivityTime(activity.timestamp)}
                          </span>
                          <p class="text-sm">
                            <span class="font-medium text-[#6b48ff]"
                              >{activity.user}</span
                            >
                            {activity.action}
                            <span class="font-medium">{activity.item}</span>
                            {#if activity.targetStatus}
                              to <span class="font-medium"
                                >{activity.targetStatus}</span
                              >
                            {/if}
                          </p>
                        </div>
                      </div>
                    </div>
                  {/each}
                </div>
              </div>
            </div>
          {/each}
        </div>
      </div>
    </div>
  {/if}
</main>

<!-- Task Creation Modal -->
{#if showTaskModal}
  <div
    class="fixed inset-0 bg-black z-50 flex items-center justify-center modal"
  >
    <div
      class="bg-white w-full max-w-2xl max-h-[90vh] overflow-y-auto shadow-lg"
    >
      <div class="p-4 border-b border-gray-100">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-semibold">Create New Tasks</h3>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeTaskModal}
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
      </div>

      <div class="p-4">
        <!-- Task form repeater -->
        {#each newTasks as task, index}
          <div class="card p-3 mb-4 bg-white shadow-sm">
            <div class="flex justify-between items-center mb-2">
              <h4 class="font-medium">Task {index + 1}</h4>
              {#if newTasks.length > 1}
                <button
                  class="text-gray-500 hover:text-red-500"
                  onclick={() => removeTask(index)}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5"
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
              {/if}
            </div>

            <div class="space-y-3">
              <div>
                <label class="block text-sm font-medium mb-1">Task Name*</label>
                <input
                  type="text"
                  class="w-full p-2 border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  placeholder="Enter task name"
                  bind:value={task.name}
                />
              </div>

              <div>
                <label class="block text-sm font-medium mb-1"
                  >Description*</label
                >
                <textarea
                  class="w-full p-2 bg-white border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  placeholder="Enter task description"
                  rows="2"
                  bind:value={task.description}
                ></textarea>
              </div>

              <div>
                <label class="block text-sm font-medium mb-1">Note</label>
                <textarea
                  class="w-full p-2 border bg-white border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  placeholder="Note if there is"
                  rows="1"
                  bind:value={task.note}
                ></textarea>
              </div>

              <div class="grid grid-cols-2 gap-3">
                <div>
                  <label class="block text-sm font-medium mb-1"
                    >Assigned To</label
                  >
                  <select
                    class="w-full p-2 border bg-white border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                    bind:value={task.assignedTo}
                  >
                    <option value="">Unassigned</option>
                    {#each teamMembers as member}
                      <option value={member.id}>{member.name}</option>
                    {/each}
                  </select>
                </div>
              </div>
            </div>
          </div>
        {/each}

        <!-- Add another task button -->
        <div class="flex justify-center">
          <button
            class="p-2 card shadow-sm bg-white text-gray-600 hover:bg-gray-50 mb-4 flex items-center justify-center"
            onclick={addAnotherTask}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5 mr-1 text-[#6b48ff]"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 4v16m8-8H4"
              />
            </svg>
            Add Another Task
          </button>
        </div>
      </div>

      <div class="p-4 border-t border-gray-100 flex justify-end space-x-3">
        <button
          class="px-4 py-2 border border-gray-200 rounded-lg hover:bg-gray-50"
          onclick={closeTaskModal}
        >
          Cancel
        </button>
        <button
          class="px-4 py-2 bg-[#6b48ff] text-white rounded-lg hover:bg-[#5a3dd4]"
          onclick={createTasksWithConfirm}
        >
          Create Tasks
        </button>
      </div>
    </div>
  </div>
{/if}

<!-- Team Members Modal -->
{#if showTeamMembersModal}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center modal"
  >
    <div
      class="bg-white rounded-lg w-full max-w-2xl max-h-[90vh] flex flex-col overflow-hidden shadow-lg"
    >
      <div class="p-4 border-b border-gray-100 bg-white flex-shrink-0">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-semibold">Team Members</h3>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeTeamMembersModal}
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
      </div>

      <!-- Content -->
      <div class="p-4 pb-6 overflow-y-auto flex-grow">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          {#each teamMembers as member}
            <div class="card p-3 bg-white shadow-sm">
              <div class="flex items-center space-x-3">
                <div
                  class="w-12 h-12 rounded-full bg-gray-200 flex items-center justify-center text-lg font-medium text-gray-600"
                >
                  {member.name
                    .split(" ")
                    .map((name) => name[0])
                    .join("")}
                </div>
                <div>
                  <p class="text-sm font-medium">{member.name}</p>
                  <p class="text-xs text-gray-500">{member.role}</p>
                  <p class="text-xs text-gray-500">{member.email}</p>
                </div>
              </div>
            </div>
          {/each}
        </div>
      </div>

      <div class="p-4 border-t border-gray-100 flex justify-end">
        <button
          class="px-4 py-2 border border-gray-200 rounded-lg hover:bg-gray-50"
          onclick={closeTeamMembersModal}
        >
          Close
        </button>
      </div>
    </div>
  </div>
{/if}

<!-- Edit Task Modal -->
{#if showEditTaskModal && currentEditingTask}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center modal"
  >
    <div
      class="bg-white rounded-lg w-full max-w-md max-h-[90vh] flex flex-col overflow-hidden shadow-lg"
    >
      <div class="p-4 border-b border-gray-100 bg-white flex-shrink-0">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-semibold">Edit Task</h3>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeEditTaskModal}
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
      </div>

      <div class="p-4 overflow-y-auto flex-grow">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1">Task Name*</label>
            <input
              type="text"
              class="w-full p-2 border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="Enter task name"
              bind:value={currentEditingTask.title}
            />
          </div>

          <div>
            <label class="block text-sm font-medium mb-1">Description</label>
            <textarea
              class="w-full p-2 border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="Enter task description"
              rows="3"
              bind:value={currentEditingTask.description}
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium mb-1">Note</label>
            <textarea
              class="w-full p-2 border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="Additional notes"
              rows="2"
              bind:value={currentEditingTask.note}
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium mb-1">Status</label>
            <select
              class="w-full p-2 border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              bind:value={currentEditingTask.status}
            >
              <option value="todo">To-do</option>
              <option value="in_progress">In Progress</option>
              <option value="review">Review</option>
              <option value="done">Done</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium mb-1">Assigned To</label>
            <select
              class="w-full p-2 border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              bind:value={currentEditingTask.assigned_to}
            >
              <option value="">Unassigned</option>
              {#each teamMembers as member}
                <option value={member.id}>{member.name}</option>
              {/each}
            </select>
          </div>
        </div>
      </div>

      <div class="p-4 border-t border-gray-100 flex justify-between">
        <button
          class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600"
          onclick={deleteTask}
        >
          Delete
        </button>

        <div class="flex space-x-3">
          <button
            class="px-4 py-2 border border-gray-200 rounded-lg hover:bg-gray-50"
            onclick={closeEditTaskModal}
          >
            Cancel
          </button>
          <button
            class="px-4 py-2 bg-[#6b48ff] text-white rounded-lg hover:bg-[#5a3dd4]"
            onclick={updateTask}
          >
            Save Changes
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}

<style>
  .modal {
    background-color: rgba(0, 0, 0, 0.5);
  }

  .card {
    border-radius: 0.5rem;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  }
</style>
