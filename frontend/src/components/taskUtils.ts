// Các hàm format date và time
function formatDate(dateString: any) {
  if (!dateString) return "N/A";

  // Nếu là timestamp số (milliseconds)
  if (typeof dateString === "number") {
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  }

  // Nếu là ISO string hoặc string khác
  try {
    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "Invalid date";

    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  } catch (e) {
    console.error("Error formatting date:", e);
    return "Invalid date";
  }
}

// Format date for activities grouping
function formatActivityDate(dateString: any) {
  if (!dateString) return "Unknown date";

  try {
    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "Invalid date";

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
  } catch (e) {
    console.error("Error formatting activity date:", e);
    return "Unknown date";
  }
}

// Format time
function formatActivityTime(dateString: any) {
  if (!dateString) return "";

  try {
    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "";

    return date.toLocaleTimeString("en-US", {
      hour: "2-digit",
      minute: "2-digit",
    });
  } catch (e) {
    console.error("Error formatting activity time:", e);
    return "";
  }
}

// Standardize task status between frontend and backend
function normalizeTaskStatus(status: any) {
  if (!status) return "todo";

  const lowercaseStatus = status.toLowerCase().trim();

  // Map various formats to standardized statuses
  switch (lowercaseStatus) {
    case "todo":
    case "to-do":
    case "to do":
    case "backlog":
      return "todo";

    case "in_progress":
    case "inprogress":
    case "in progress":
    case "in-progress":
    case "ongoing":
    case "started":
      return "in_progress";

    case "review":
    case "reviewing":
    case "in review":
    case "for review":
    case "testing":
      return "review";

    case "done":
    case "completed":
    case "complete":
    case "finished":
      return "done";

    default:
      console.warn(`Unknown status: ${status}, defaulting to 'todo'`);
      return "todo";
  }
}

// Function to process tasks from WebSocket
function processTasksFromServer(tasksFromServer: any) {
  if (!Array.isArray(tasksFromServer)) {
    console.error("Invalid tasks data:", tasksFromServer);
    return [];
  }

  return tasksFromServer.map((task) => ({
    id: task.id || task._id || "",
    title: task.title || task.name || "Untitled Task",
    description: task.description || "",
    status: normalizeTaskStatus(task.status),
    assigned_to: task.assigned_to || task.assignedTo || "",
    note: task.note || "",
  }));
}

// Function to process activities from WebSocket
function processActivitiesFromServer(
  activitiesFromServer: any,
  userMap?: Record<string, string>
) {
  if (!Array.isArray(activitiesFromServer)) {
    console.error("Invalid activities data:", activitiesFromServer);
    return [];
  }

  // Map activities with proper structure
  const mappedActivities = activitiesFromServer.map((activity) => {
    // Determine action based on type
    let action = "updated";
    switch (activity.type) {
      case "create":
        action = "created";
        break;
      case "update":
        action = "updated";
        break;
      case "delete":
        action = "deleted";
        break;
      case "move":
        action = "moved";
        break;
      default:
        action = activity.type || "updated";
    }

    // Try to map user ID to username if userMap is provided
    let userName = activity.done_by;
    if (userMap && activity.user_id && userMap[activity.user_id]) {
      userName = userMap[activity.user_id];
    }

    return {
      id:
        activity.id || activity._id || Math.random().toString(36).substr(2, 9),
      user: userName,
      action: action,
      item: activity.title || activity.item || "Unknown Item",
      targetStatus: activity.to || activity.targetStatus || "",
      timestamp:
        activity.created_at || activity.timestamp || new Date().toISOString(),
    };
  });

  // Sort activities by timestamp (newest first)
  return mappedActivities.sort((a, b) => {
    return new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime();
  });
}

// Handle tasks movement between columns
function handleTaskMove(
  taskId: any,
  fromStatus: any,
  toStatus: any,
  socket: any
) {
  if (!socket) {
    console.error("Socket not connected");
    return;
  }

  console.log(`Moving task: ${taskId} from ${fromStatus} to ${toStatus}`);

  // Send the update to server
  try {
    socket.emit("message", {
      type: "update",
      content: {
        taskId: taskId,
        status: toStatus,
      },
    });
  } catch (error) {
    console.error("Error sending task movement to server:", error);
  }
}

// Initialize SortableJS with proper options
async function initSortable(
  getTasksByStatus: any,
  handleTaskMove: any,
  socket: any
) {
  if (typeof window === "undefined") return [];

  try {
    // Lazy load Sortable
    const SortableModule = await import("sortablejs");
    const Sortable = SortableModule.default;

    const sortables: any[] = [];

    // Destroy old instances if any
    document.querySelectorAll(".kanban-column").forEach((column) => {
      const existingSortable = Sortable.get(column as HTMLElement);
      if (existingSortable) {
        existingSortable.destroy();
      }
    });

    // Create new instances
    document.querySelectorAll(".kanban-column").forEach((column) => {
      const status = column.getAttribute("data-status");
      if (!status) {
        console.error("Missing data-status attribute on column", column);
        return;
      }

      console.log(`Initializing sortable for status: ${status}`);

      const sortable = new Sortable(column as HTMLElement, {
        group: "tasks",
        animation: 150,
        ghostClass: "task-ghost",
        chosenClass: "task-chosen",
        dragClass: "task-drag",
        filter: ".ignore-elements", // Elements that should not be draggable
        removeCloneOnHide: true, // Important: prevent ghost clones
        forceFallback: true, // Use custom drag ghost
        fallbackClass: "sortable-fallback",
        onEnd: function (evt) {
          const taskId = evt.item.getAttribute("data-id");
          const fromStatus = evt.from.getAttribute("data-status");
          const toStatus = evt.to.getAttribute("data-status");

          // Validate all required data is present
          if (!taskId || !fromStatus || !toStatus) {
            console.error("Missing data for task movement", {
              taskId,
              fromStatus,
              toStatus,
            });
            return;
          }

          // Prevent duplicate cards by removing the dragged element
          // This prevents the UI from having duplicates until the data is refreshed
          if (fromStatus !== toStatus) {
            // Remove the dragged item from DOM immediately
            evt.item.remove();

            // Then handle the task move (which will update state and refresh UI)
            handleTaskMove(taskId, fromStatus, toStatus, socket);
          }
        },
      });

      sortables.push(sortable);
    });

    return sortables;
  } catch (error) {
    console.error("Error initializing Sortable:", error);
    return [];
  }
}

// Setup WebSocket connection
function setupWebSocket(projectId: any, userId: any, onDataReceived: any) {
  if (typeof window === "undefined") return null;

  try {
    // Get the base URL from the current window location
    const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
    const host = window.location.host; // This will include port if present

    // Determine the correct WebSocket URL based on environment
    let wsUrl;

    // Check if we're in production (deployed to skillforge.ikniz.id.vn)
    if (host.includes("skillforge.ikniz.id.vn")) {
      // For production: Use the same host but with WS protocol
      wsUrl = `${protocol}//${host}/ws/task/${projectId}/${userId}`;
    } else {
      // For development: Use explicitly defined backend server (likely on port 8080)
      const devHost = window.location.hostname;
      wsUrl = `${protocol}//${devHost}:8080/ws/task/${projectId}/${userId}`;
    }

    console.log("Connecting with WebSocket to:", wsUrl);

    const ws = new WebSocket(wsUrl);

    ws.onopen = () => {
      console.log("WebSocket connection established");
    };

    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        console.log("Received via WebSocket:", data);

        // Process and standardize the data
        const processedData = {
          tasks: data.tasks ? processTasksFromServer(data.tasks) : [],
          activities: data.activities
            ? processActivitiesFromServer(data.activities)
            : [],
        };

        // Call the callback with processed data
        if (onDataReceived && typeof onDataReceived === "function") {
          onDataReceived(processedData);
        }
      } catch (e) {
        console.error("Error parsing WebSocket data:", e);
      }
    };

    ws.onerror = (err) => {
      console.error("WebSocket error:", err);
    };

    ws.onclose = () => {
      console.log("WebSocket connection closed");
    };

    // Create a Socket.io-like interface
    const socketInterface = {
      emit: (eventName: any, data: any) => {
        if (ws.readyState === WebSocket.OPEN) {
          ws.send(
            JSON.stringify({
              type: data.type,
              content: data.content,
            })
          );
        } else {
          console.error("WebSocket is not connected");
        }
      },
      disconnect: () => {
        ws.close();
      },
    };

    return socketInterface;
  } catch (error) {
    console.error("Error setting up WebSocket:", error);
    return null;
  }
}

// Task CRUD functions
function createTask(socket: any, taskInput: any) {
  if (!socket) {
    console.error("Socket not connected");
    return;
  }

  const message = {
    type: "create",
    content: {
      title: taskInput.name || taskInput.title,
      description: taskInput.description || "",
      note: taskInput.note || "",
      assigned_to: taskInput.assignedTo || taskInput.assigned_to || "",
      status: normalizeTaskStatus(taskInput.status) || "todo",
    },
  };

  console.log("Sending create task message:", message);

  try {
    socket.emit("message", message);
    console.log("Task creation message sent successfully");
    return true;
  } catch (error) {
    console.error("Error sending task creation message:", error);
    return false;
  }
}

function updateTask(socket: any, task: any) {
  if (!socket || !task || !task.id) {
    console.error("Socket not connected or invalid task data", task);
    return;
  }

  const message = {
    type: "update",
    content: {
      taskId: task.id,
      title: task.title,
      description: task.description || "",
      note: task.note || "",
      status: normalizeTaskStatus(task.status),
      assigned_to: task.assigned_to || "",
    },
  };

  console.log("Sending update task message:", message);

  try {
    socket.emit("message", message);
    console.log("Task update message sent successfully");
    return true;
  } catch (error) {
    console.error("Error sending task update message:", error);
    return false;
  }
}

function deleteTask(socket: any, taskId: any) {
  if (!socket || !taskId) {
    console.error("Socket not connected or invalid task ID");
    return;
  }

  const message = {
    type: "delete",
    content: taskId,
  };

  console.log("Sending delete task message:", message);

  try {
    socket.emit("message", message);
    console.log("Task deletion message sent successfully");
    return true;
  } catch (error) {
    console.error("Error sending task deletion message:", error);
    return false;
  }
}

// Export the functions for use in Svelte component
export {
  formatDate,
  formatActivityDate,
  formatActivityTime,
  normalizeTaskStatus,
  processTasksFromServer,
  processActivitiesFromServer,
  handleTaskMove,
  initSortable,
  setupWebSocket,
  createTask,
  updateTask,
  deleteTask,
};
