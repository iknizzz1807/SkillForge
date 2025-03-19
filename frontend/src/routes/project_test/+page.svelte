<script lang="ts">
  import { enhance } from "$app/forms";
  import type { ActionData, PageData } from "./$types";

  let { form, data }: { form: ActionData; data: PageData } = $props();

  let creating: boolean = $state(false);

  $effect(() => {
    if (form?.project) {
      projectsDisplay.push(form.project);
      form.project = undefined;
      skills = [];
    }
  });

  type ProjectDisplay = {
    id: string;
    title: string;
    description: string;
    skills: string[];
    timeline: string;
    created_by: string;
    status: string;
    repo_url: string;
    created_at: string;
  };

  let projectsDisplay: ProjectDisplay[] = $state(data.projects);
  let errorLoadingProjects: string | null = $state(data.error);

  // Bind value input
  let title: string = $state("");
  let description: string = $state("");
  let skills: string[] = $state([]);
  let timeline: string = $state("");

  // Categories for projects
  const skillsAvailable = ["Python", "C++", "JavaScript", "Go", "HTML", "CSS"];

  function toggleSkill(skill: string) {
    if (skills.includes(skill)) {
      skills = skills.filter((s) => s !== skill);
    } else {
      skills = [...skills, skill];
    }
  }

  // Get status badge class
  function getStatusClass(status: string): string {
    switch (status.toLowerCase()) {
      case "open":
        return "status-open";
      case "in progress":
        return "status-progress";
      case "completed":
        return "status-completed";
      default:
        return "status-default";
    }
  }

  // Chuyển đổi ngày thành format thân thiện
  function formatDate(dateString: string): string {
    const date = new Date(dateString);
    return new Intl.DateTimeFormat("en-US", {
      year: "numeric",
      month: "short",
      day: "numeric",
    }).format(date);
  }
</script>

<form method="POST" use:enhance>
  <div class="form-group">
    <label for="title">Project Title</label>
    <input type="text" id="title" name="title" bind:value={title} required />
  </div>

  <div class="form-group">
    <label for="description">Description</label>
    <textarea
      id="description"
      name="description"
      bind:value={description}
      required
    ></textarea>
  </div>

  <!-- Multi-select skills input -->
  <div class="form-group">
    <label for="skills">Skills Required</label>
    <div class="skills-selector">
      {#each skillsAvailable as skill}
        <div class="skill-item">
          <label class="skill-label">
            <input
              type="checkbox"
              checked={skills.includes(skill)}
              onchange={() => toggleSkill(skill)}
            />
            <span class="skill-name">{skill}</span>
          </label>
        </div>
      {/each}
    </div>

    <!-- Hidden inputs to send selected skills -->
    {#each skills as skill}
      <input type="hidden" name="skills" value={skill} />
    {/each}

    <div class="selected-skills">
      <p>
        Selected skills:
        {#if skills.length === 0}
          <span class="none-selected">None</span>
        {:else}
          {#each skills as skill, i}
            <span class="skill-tag">
              {skill}
              <button
                type="button"
                class="remove-skill"
                onclick={() => toggleSkill(skill)}>×</button
              >
            </span>
          {/each}
        {/if}
      </p>
    </div>
  </div>

  <div class="form-group">
    <label for="timeline">Timeline (weeks)</label>
    <input
      type="number"
      id="timeline"
      name="timeline"
      bind:value={timeline}
      required
    />
  </div>

  <button type="submit" class="submit-btn">Create Project</button>

  {#if form?.error}
    <div class="error-message">
      {form.error}
    </div>
  {/if}
</form>

<!-- Display projects -->
<section class="projects-section">
  <h2>Available Projects</h2>

  {#if errorLoadingProjects}
    <div class="error-message">
      {errorLoadingProjects}
    </div>
  {:else if projectsDisplay.length === 0}
    <div class="no-projects">
      <p>No projects available at the moment.</p>
    </div>
  {:else}
    <div class="projects-grid">
      {#each projectsDisplay as project (project.id)}
        <div class="project-card">
          <div class="project-header">
            <h3 class="project-title">{project.title}</h3>
            <span class="status-badge {getStatusClass(project.status)}">
              {project.status}
            </span>
          </div>

          <p class="project-description">{project.description}</p>

          <div class="project-skills">
            {#each project.skills as skill}
              <span class="project-skill-tag">{skill}</span>
            {/each}
          </div>

          <div class="project-meta">
            <div class="meta-item">
              <span class="meta-label">Timeline:</span>
              <span class="meta-value">{project.timeline} weeks</span>
            </div>
            <div class="meta-item">
              <span class="meta-label">Created by:</span>
              <span class="meta-value">{project.created_by}</span>
            </div>
            <div class="meta-item">
              <span class="meta-label">Posted:</span>
              <span class="meta-value">{formatDate(project.created_at)}</span>
            </div>
          </div>

          {#if project.repo_url}
            <a
              href={project.repo_url}
              target="_blank"
              rel="noopener noreferrer"
              class="repo-link"
            >
              View Repository
            </a>
          {/if}

          <a href={`/projects/${project.id}`} class="view-details">
            View Details
          </a>
        </div>
      {/each}
    </div>
  {/if}
</section>

<style>
  .skills-selector {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    margin-bottom: 10px;
  }

  .skill-item {
    background-color: #f5f5f5;
    border-radius: 4px;
    padding: 6px 12px;
  }

  .skill-label {
    display: flex;
    align-items: center;
    gap: 6px;
    cursor: pointer;
  }

  .selected-skills {
    margin-top: 10px;
  }

  .skill-tag {
    display: inline-flex;
    align-items: center;
    background-color: #e0f7fa;
    color: #00796b;
    border-radius: 16px;
    padding: 4px 10px;
    margin-right: 5px;
    margin-bottom: 5px;
    font-size: 14px;
  }

  .remove-skill {
    background: none;
    border: none;
    color: #00796b;
    margin-left: 4px;
    cursor: pointer;
    font-size: 16px;
    font-weight: bold;
    padding: 0 4px;
  }

  .none-selected {
    color: #999;
    font-style: italic;
  }

  .form-group {
    margin-bottom: 20px;
  }

  input[type="text"],
  input[type="number"],
  textarea {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  textarea {
    min-height: 100px;
  }

  .submit-btn {
    background-color: #0277bd;
    color: white;
    border: none;
    padding: 12px 20px;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
  }

  .submit-btn:hover {
    background-color: #01579b;
  }

  .error-message {
    color: #d32f2f;
    margin-top: 15px;
    padding: 10px;
    background-color: #ffebee;
    border-radius: 4px;
  }

  /* Projects Section Styles */
  .projects-section {
    margin-top: 40px;
    padding-top: 30px;
    border-top: 1px solid #eee;
  }

  .projects-section h2 {
    margin-bottom: 20px;
    font-size: 24px;
    color: #333;
  }

  .no-projects {
    padding: 20px;
    background-color: #f5f5f5;
    border-radius: 4px;
    text-align: center;
  }

  .projects-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 20px;
  }

  .project-card {
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    padding: 20px;
    display: flex;
    flex-direction: column;
    transition:
      transform 0.2s,
      box-shadow 0.2s;
  }

  .project-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 15px rgba(0, 0, 0, 0.1);
  }

  .project-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    margin-bottom: 12px;
  }

  .project-title {
    font-size: 18px;
    margin: 0;
    color: #333;
  }

  .status-badge {
    padding: 4px 8px;
    border-radius: 12px;
    font-size: 12px;
    font-weight: bold;
    text-transform: uppercase;
  }

  .status-open {
    background-color: #e3f2fd;
    color: #1565c0;
  }

  .status-progress {
    background-color: #fff8e1;
    color: #ff8f00;
  }

  .status-completed {
    background-color: #e8f5e9;
    color: #2e7d32;
  }

  .status-default {
    background-color: #f5f5f5;
    color: #616161;
  }

  .project-description {
    color: #666;
    margin-bottom: 15px;
    font-size: 14px;
    line-height: 1.5;
    flex-grow: 1;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .project-skills {
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
    margin-bottom: 15px;
  }

  .project-skill-tag {
    background-color: #f1f8e9;
    color: #558b2f;
    padding: 4px 8px;
    border-radius: 12px;
    font-size: 12px;
  }

  .project-meta {
    margin-bottom: 15px;
    font-size: 13px;
  }

  .meta-item {
    margin-bottom: 6px;
    display: flex;
  }

  .meta-label {
    color: #757575;
    width: 80px;
  }

  .meta-value {
    color: #424242;
    font-weight: 500;
  }

  .repo-link,
  .view-details {
    display: inline-block;
    margin-top: 10px;
    padding: 6px 12px;
    text-align: center;
    text-decoration: none;
    border-radius: 4px;
    font-size: 14px;
    font-weight: 500;
    transition: background-color 0.2s;
  }

  .repo-link {
    background-color: #f5f5f5;
    color: #424242;
    margin-right: 8px;
  }

  .repo-link:hover {
    background-color: #e0e0e0;
  }

  .view-details {
    background-color: #e0f7fa;
    color: #00796b;
  }

  .view-details:hover {
    background-color: #b2ebf2;
  }
</style>
