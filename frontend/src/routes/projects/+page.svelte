<script lang="ts">
  import { enhance } from "$app/forms";
  import type { ActionData, PageData } from "./$types";

  let { form, data }: { form: ActionData; data: PageData } = $props();

  type Project = {
    title: string;
    description: string;
    category: string;
    timeline: string;
  };

  let projects: Project[] = $state(data.projects);
  let errorLoadingProjects: string | null = $state(data.error);

  let creating: boolean = $state(false);
  let loading: boolean = $state(false);

  // Bind value input
  let title: string = $state("");
  let description: string = $state("");
  let category: string = $state("");
  let timeline: string = $state("");

  // Categories for projects
  const categories = [
    "Web Development",
    "Mobile Development",
    "Machine Learning",
    "Data Science",
    "Cloud Computing",
    "DevOps",
    "UI/UX Design",
    "Blockchain",
    "Cybersecurity",
    "Game Development",
  ];

  // Difficulty levels
  const difficultyLevels = [
    { value: "BEGINNER", label: "Beginner" },
    { value: "INTERMEDIATE", label: "Intermediate" },
    { value: "ADVANCED", label: "Advanced" },
  ];

  // Toggle project creation form
  function toggleCreateProject() {
    creating = !creating;
  }

  // Handle form submission with enhance
  function handleSubmit() {
    loading = true;
    // return async ({ result, update }) => {
    //   loading = false;
    //   if (result.type === "success") {
    //     creating = false;
    //     await invalidateAll(); // Refresh data
    //   }
    //   update();
    // };
  }
</script>

<div class="projects-container">
  <div class="projects-header">
    <h1>Projects</h1>
    <button class="create-btn" onclick={toggleCreateProject}>
      {creating ? "Cancel" : "Create Project"}
    </button>
  </div>

  {#if creating}
    <div class="create-project-form">
      <h2>Create New Project</h2>

      {#if form?.error}
        <div class="error-message">
          {form.message}
        </div>
      {/if}

      <form method="POST" action="?/createProject" use:enhance={handleSubmit}>
        <div class="form-group">
          <label for="title"
            >Project Title <span class="required">*</span></label
          >
          <input
            type="text"
            id="title"
            name="title"
            required
            bind:value={title}
            placeholder="Enter project title"
          />
        </div>

        <div class="form-group">
          <label for="description"
            >Description <span class="required">*</span></label
          >
          <textarea
            id="description"
            name="description"
            required
            bind:value={description}
            placeholder="Describe your project"
            rows="4"
          ></textarea>
        </div>

        <div class="form-row">
          <div class="form-group half">
            <label for="category"
              >Category <span class="required">*</span></label
            >
            <select
              id="category"
              name="category"
              required
              bind:value={category}
            >
              <option value="" disabled selected>Select a category</option>
              {#each categories as category}
                <option value={category}>
                  {category}
                </option>
              {/each}
            </select>
          </div>
        </div>

        <!-- <div class="form-group">
          <label for="requirements"
            >Requirements <span class="required">*</span></label
          >
          <textarea
            id="requirements"
            name="requirements"
            required
            value={form?.requirements || ""}
            placeholder="List the requirements for this project (e.g. technologies, skills)"
            rows="3"
          ></textarea>
        </div> -->

        <!-- <div class="form-group">
          <label for="learningOutcomes"
            >Learning Outcomes <span class="required">*</span></label
          >
          <textarea
            id="learningOutcomes"
            name="learningOutcomes"
            required
            value={form?.learningOutcomes || ""}
            placeholder="What will users learn from this project?"
            rows="3"
          ></textarea>
        </div> -->

        <div class="form-group">
          <label for="estimatedTime"
            >Estimated Completion Time (hours) <span class="required">*</span
            ></label
          >
          <input
            type="number"
            id="estimatedTime"
            name="estimatedTime"
            required
            min="1"
            bind:value={timeline}
            placeholder="e.g. 10"
          />
        </div>

        <div class="form-actions">
          <button type="submit" class="submit-btn" disabled={loading}>
            {loading ? "Creating..." : "Create Project"}
          </button>
          <button
            type="button"
            class="cancel-btn"
            onclick={toggleCreateProject}
            disabled={loading}
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  {:else if projects.length}
    <div class="projects-list">
      <!-- Project list would go here -->
      <p>Your projects will be displayed here</p>
    </div>
  {:else if errorLoadingProjects}
    <div class="error-message">
      {errorLoadingProjects}
    </div>
  {:else}
    <div class="no-projects">
      <p>You haven't created any projects yet.</p>
      <button class="create-btn" onclick={toggleCreateProject}
        >Create Your First Project</button
      >
    </div>
  {/if}
</div>

<style>
  .projects-container {
    max-width: 1000px;
    margin: 0 auto;
    padding: 2rem 1rem;
  }

  .projects-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
  }

  h1 {
    margin: 0;
    font-size: 2rem;
  }

  h2 {
    margin-top: 0;
    margin-bottom: 1.5rem;
    font-size: 1.5rem;
  }

  .create-btn {
    background-color: #4a6cf7;
    color: white;
    border: none;
    border-radius: 4px;
    padding: 0.5rem 1rem;
    cursor: pointer;
    font-weight: 500;
  }

  .create-btn:hover {
    background-color: #3a5cd7;
  }

  .create-project-form {
    background-color: #f9f9f9;
    border-radius: 8px;
    padding: 2rem;
    margin-bottom: 2rem;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }

  .form-group {
    margin-bottom: 1.25rem;
  }

  .form-row {
    display: flex;
    gap: 1rem;
    margin-bottom: 1.25rem;
  }

  .half {
    flex: 1;
  }

  label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
  }

  .required {
    color: #e53935;
  }

  input,
  select,
  textarea {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 1rem;
    font-family: inherit;
  }

  input:focus,
  select:focus,
  textarea:focus {
    outline: none;
    border-color: #4a6cf7;
    box-shadow: 0 0 0 2px rgba(74, 108, 247, 0.2);
  }

  .form-actions {
    display: flex;
    gap: 1rem;
    margin-top: 1.5rem;
  }

  .submit-btn,
  .cancel-btn {
    padding: 0.75rem 1.5rem;
    border-radius: 4px;
    cursor: pointer;
    font-weight: 500;
    font-size: 1rem;
  }

  .submit-btn {
    background-color: #4a6cf7;
    color: white;
    border: none;
  }

  .submit-btn:hover:not(:disabled) {
    background-color: #3a5cd7;
  }

  .cancel-btn {
    background-color: transparent;
    border: 1px solid #ddd;
    color: #666;
  }

  .cancel-btn:hover:not(:disabled) {
    background-color: #f2f2f2;
  }

  button:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .error-message {
    background-color: #ffebee;
    color: #c62828;
    padding: 0.75rem;
    border-radius: 4px;
    margin-bottom: 1.5rem;
  }

  .no-projects {
    text-align: center;
    padding: 3rem 0;
    color: #666;
  }

  .no-projects p {
    margin-bottom: 1.5rem;
  }
</style>
