<script lang="ts">
  import { goto } from "$app/navigation";

  let member: number = $state(1);
  let description: string = $state("");
  let isSubmitting: boolean = $state(false);
  let error: string | null = $state(null);

  // Khai báo các biến để lưu trữ dữ liệu form
  let title: string = $state("");
  let skills: string[] = $state([]);
  let startDate: string = $state("");
  let endDate: string = $state("");

  // Thêm state cho modal thành công
  let showSuccessModal: boolean = $state(false);
  let createdProjectInfo: any = $state(null);

  async function handleSubmit(event: Event) {
    event.preventDefault();
    isSubmitting = true;
    error = null;

    try {
      // Lấy dữ liệu từ form
      const formElement = event.target as HTMLFormElement;
      const formSkills = Array.from(
        formElement.querySelector("#skills") as HTMLSelectElement
      )
        .filter((option) => (option as HTMLOptionElement).selected)
        .map((option) => (option as HTMLOptionElement).value);

      // Tạo payload cho API
      const payload = {
        title: title,
        description: description,
        skills: formSkills,
        max_member: member,
        start_time: startDate,
        end_time: endDate,
      };

      // Gửi request đến API endpoint
      const response = await fetch("/api/projects", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "Failed to create project");
      }

      // Xử lý phản hồi thành công
      const createdProject = await response.json();
      console.log("Project created:", createdProject);

      // Lưu thông tin project đã tạo và hiển thị modal thành công
      createdProjectInfo = createdProject;
      showSuccessModal = true;
    } catch (err: any) {
      error = err.message || "An unexpected error occurred";
      console.error("Error creating project:", err);
    } finally {
      isSubmitting = false;
    }
  }

  // Hàm đóng modal và chuyển hướng về trang projects
  function closeModalAndRedirect() {
    showSuccessModal = false;
    goto("/project");
  }

  // Hàm đóng modal và giữ nguyên trang để tạo project mới
  function closeModalAndStay() {
    showSuccessModal = false;
    // Reset form
    title = "";
    description = "";
    skills = [];
    startDate = "";
    endDate = "";
    member = 1;

    // Reset select element (cần phải reset DOM element vì Svelte không tự động cập nhật select multiple)
    const selectElement = document.querySelector(
      "#skills"
    ) as HTMLSelectElement;
    if (selectElement) {
      Array.from(selectElement.options).forEach((option) => {
        option.selected = false;
      });
    }
  }

  // Xử lý khi người dùng chọn skills
  function handleSkillsChange(event: Event) {
    const select = event.target as HTMLSelectElement;
    skills = Array.from(select.selectedOptions).map((option) => option.value);
  }
</script>

<svelte:head>
  <title>Create New Project</title>
</svelte:head>

<main class="flex-1 pr-4 pl-4 ml-64 pt-4">
  <div class="flex items-center mb-6">
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
    <h2 class="text-xl font-semibold">Create New Project</h2>
  </div>

  {#if error}
    <div
      class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4"
    >
      <p>{error}</p>
    </div>
  {/if}

  <div class="bg-white rounded-lg shadow p-6 max-w-4xl mx-auto">
    <form class="space-y-6" onsubmit={handleSubmit}>
      <!-- Basic Project Info -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1" for="title"
              >Project Title*</label
            >
            <input
              type="text"
              id="title"
              bind:value={title}
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="Enter a descriptive title"
              required
            />
          </div>

          <div>
            <label for="skills" class="block text-sm font-medium mb-1"
              >Required Skills*</label
            >
            <select
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              multiple
              size="4"
              id="skills"
              onchange={handleSkillsChange}
              required
            >
              <option>React</option>
              <option>Node.js</option>
              <option>Python</option>
              <option>MongoDB</option>
              <option>UI/UX Design</option>
              <option>SwiftUI</option>
              <option>DevOps</option>
              <option>AWS</option>
              <option>Machine Learning</option>
              <option>Vue.js</option>
            </select>
            <p class="text-xs text-gray-500 mt-1">
              Hold Ctrl/Cmd to select multiple
            </p>
          </div>

          <div>
            <label for="max-member" class="block text-sm font-medium mb-1"
              >Team Size*</label
            >
            <div class="flex items-center space-x-2">
              <input
                id="max-member"
                type="range"
                min="1"
                max="20"
                bind:value={member}
                class="w-full accent-[#6b48ff]"
              />
              <span class="text-sm w-8 text-center">{member}</span>
            </div>
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1</span>
              <span>20</span>
            </div>
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1" for=""
              >Project Timeline*</label
            >
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="text-xs text-gray-500 block mb-1" for="start-date"
                  >Start Date</label
                >
                <input
                  type="date"
                  id="start-time"
                  bind:value={startDate}
                  class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  required
                />
              </div>
              <div>
                <label class="text-xs text-gray-500 block mb-1" for="end-time"
                  >End Date</label
                >
                <input
                  type="date"
                  id="end-time"
                  bind:value={endDate}
                  class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  required
                />
              </div>
            </div>
          </div>

          <!-- External Links -->
          <div>
            <label class="block text-sm font-medium mb-1" for=""
              >External References (Optional)</label
            >
            <div class="space-y-3">
              <div class="flex space-x-2">
                <input
                  type="text"
                  placeholder="GitHub Repository URL"
                  class="flex-1 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                />
                <select
                  class="p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                >
                  <option>GitHub</option>
                  <option>Figma</option>
                  <option>Documentation</option>
                  <option>Other</option>
                </select>
              </div>
              <button
                type="button"
                class="text-sm text-[#6b48ff] flex items-center"
              >
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
                Add Another Link
              </button>
            </div>
          </div>

          <!-- Files & References -->
          <div>
            <label class="block text-sm font-medium mb-1">Attach Files</label>
            <div
              class="border border-dashed border-gray-300 rounded p-4 text-center bg-gray-50"
            >
              <div class="space-y-1">
                <svg
                  class="mx-auto h-6 w-6 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
                  ></path>
                </svg>
                <p class="text-sm text-gray-500">
                  Drag & drop files or <span class="text-[#6b48ff] font-medium"
                    >browse</span
                  >
                </p>
                <p class="text-xs text-gray-400">Max: 10MB</p>
              </div>
              <input type="file" class="hidden" id="fileUpload" multiple />
            </div>
          </div>
        </div>
      </div>

      <!-- Project Description -->
      <div>
        <label class="block text-sm font-medium mb-1" for=""
          >Project Description*</label
        >
        <!-- Rich Text Editor -->
        <div class="border border-gray-300 rounded overflow-hidden">
          <!-- Toolbar -->
          <div class="bg-gray-100 border-b border-gray-300 p-1 flex flex-wrap">
            <button type="button" class="p-1 hover:bg-gray-200 rounded mr-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path d="M13 7H7v2h6V7z" />
                <path
                  fill-rule="evenodd"
                  d="M7 2a2 2 0 00-2 2v12a2 2 0 002 2h6a2 2 0 002-2V4a2 2 0 00-2-2H7zm6 1.5v13a.5.5 0 01-.5.5h-5a.5.5 0 01-.5-.5v-13a.5.5 0 01.5-.5h5a.5.5 0 01.5.5z"
                  clip-rule="evenodd"
                />
              </svg>
            </button>
            <button
              type="button"
              class="p-1 hover:bg-gray-200 rounded mr-1 font-bold">B</button
            >
            <button
              type="button"
              class="p-1 hover:bg-gray-200 rounded mr-1 italic">I</button
            >
            <button
              type="button"
              class="p-1 hover:bg-gray-200 rounded mr-1 underline">U</button
            >
            <span class="border-r border-gray-300 h-6 mx-1"></span>
            <button type="button" class="p-1 hover:bg-gray-200 rounded mr-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 5a1 1 0 011-1h6a1 1 0 110 2H4a1 1 0 01-1-1z"
                  clip-rule="evenodd"
                />
              </svg>
            </button>
            <button type="button" class="p-1 hover:bg-gray-200 rounded mr-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                  clip-rule="evenodd"
                />
              </svg>
            </button>
          </div>
          <!-- Editor content area with increased height -->
          <textarea
            id="description"
            bind:value={description}
            class="p-3 min-h-[250px] focus:outline-none w-full resize-none"
            placeholder="Describe your project requirements, goals, and expectations in detail..."
            required
          ></textarea>
        </div>
      </div>

      <!-- Form Actions -->
      <div class="flex justify-end space-x-3 pt-4 border-t border-gray-200">
        <button
          type="submit"
          class="bg-[#6b48ff] hover:bg-[#5a3de6] text-white py-2 px-4 rounded transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#6b48ff]"
          disabled={isSubmitting}
        >
          {isSubmitting ? "Creating..." : "Create Project"}
        </button>
      </div>
    </form>
  </div>

  <!-- Modal Thành Công -->
  {#if showSuccessModal}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 max-w-md w-full">
        <div class="text-center">
          <!-- Icon Success -->
          <div
            class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-green-100 mb-4"
          >
            <svg
              class="h-6 w-6 text-green-600"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M5 13l4 4L19 7"
              />
            </svg>
          </div>

          <h3 class="text-lg font-medium text-gray-900 mb-2">
            Project Created Successfully!
          </h3>

          <p class="text-sm text-gray-500 mb-1">
            Your project <span class="font-medium text-gray-700"
              >{createdProjectInfo?.title || "New Project"}</span
            > has been created.
          </p>

          <p class="text-sm text-gray-500 mb-6">
            Start date: <span class="font-medium"
              >{new Date(createdProjectInfo?.start_time).toLocaleDateString() ||
                "N/A"}</span
            >
            | Team size:
            <span class="font-medium"
              >{createdProjectInfo?.max_member || "N/A"}</span
            >
          </p>

          <div class="flex space-x-3 justify-center">
            <button
              class="px-4 py-2 bg-[#6b48ff] text-white rounded hover:bg-[#5a3dd3] transition-colors"
              onclick={closeModalAndRedirect}
            >
              Go to Projects
            </button>
            <button
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300 transition-colors"
              onclick={closeModalAndStay}
            >
              Create Another
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
