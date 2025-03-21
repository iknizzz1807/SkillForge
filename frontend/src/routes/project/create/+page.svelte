<script lang="ts">
  import { goto } from "$app/navigation";

  function goBack() {
    goto("/project");
  }

  let member: number = $state(1);
</script>

<svelte:head>
  <title>Create New Project</title>
</svelte:head>

<main class="flex-1 pr-4 pl-4 ml-64 pt-4">
  <div class="flex items-center mb-6">
    <button class="text-gray-500 hover:text-gray-700 mr-3" onclick={goBack}>
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
    <h2 class="text-xl font-semibold">Create New Project</h2>
  </div>

  <div class="bg-white rounded-lg shadow p-6 max-w-4xl mx-auto">
    <form class="space-y-6">
      <!-- Basic Project Info -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1">Project Title*</label>
            <input
              type="text"
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="Enter a descriptive title"
              required
            />
          </div>

          <!-- Project type
          <div>
            <label class="block text-sm font-medium mb-1">Project Type*</label>
            <div class="grid grid-cols-3 gap-3">
              <label class="cursor-pointer">
                <input
                  type="radio"
                  name="projectType"
                  value="development"
                  class="sr-only"
                />
                <div
                  class="p-3 border border-gray-300 rounded text-center hover:border-[#6b48ff] project-type-card"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-6 w-6 mx-auto mb-1"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"
                    />
                  </svg>
                  <span class="text-sm font-medium">Development</span>
                </div>
              </label>
              <label class="cursor-pointer">
                <input
                  type="radio"
                  name="projectType"
                  value="design"
                  class="sr-only"
                />
                <div
                  class="p-3 border border-gray-300 rounded text-center hover:border-[#6b48ff] project-type-card"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-6 w-6 mx-auto mb-1"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01"
                    />
                  </svg>
                  <span class="text-sm font-medium">Design</span>
                </div>
              </label>
              <label class="cursor-pointer">
                <input
                  type="radio"
                  name="projectType"
                  value="research"
                  class="sr-only"
                />
                <div
                  class="p-3 border border-gray-300 rounded text-center hover:border-[#6b48ff] project-type-card"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-6 w-6 mx-auto mb-1"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
                    />
                  </svg>
                  <span class="text-sm font-medium">Research</span>
                </div>
              </label>
            </div>
          </div> -->

          <div>
            <label class="block text-sm font-medium mb-1"
              >Required Skills*</label
            >
            <select
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              multiple
              size="4"
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
                name="max-member"
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
            <label class="block text-sm font-medium mb-1"
              >Project Timeline*</label
            >
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="text-xs text-gray-500 block mb-1"
                  >Start Date</label
                >
                <input
                  type="date"
                  class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                />
              </div>
              <div>
                <label class="text-xs text-gray-500 block mb-1">End Date</label>
                <input
                  type="date"
                  class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                />
              </div>
            </div>
          </div>

          <!-- External Links -->
          <div>
            <label class="block text-sm font-medium mb-1"
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
        <label class="block text-sm font-medium mb-1"
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
          <div
            contenteditable="true"
            class="p-3 min-h-[250px] focus:outline-none"
            placeholder="Describe your project requirements, goals, and expectations in detail..."
          ></div>
        </div>
      </div>

      <!-- Form Actions -->
      <div class="flex justify-end space-x-3 pt-4 border-t border-gray-200">
        <button
          type="button"
          class="px-4 py-2 border border-gray-300 rounded text-gray-700 hover:bg-gray-50"
          onclick={goBack}
        >
          Cancel
        </button>
        <button
          type="submit"
          class="bg-[#6b48ff] hover:bg-[#5a3de6] text-white py-2 px-4 rounded transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#6b48ff]"
        >
          Create Project
        </button>
      </div>
    </form>
  </div>
</main>

<style>
  [contenteditable]:empty:before {
    content: attr(placeholder);
    color: #9ca3af;
    display: block;
  }
</style>
