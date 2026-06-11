<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();
  let linkCopied = $state(false);

  const copyPortfolioLink = async () => {
    const url = window.location.origin + "/portfolios/" + data.id + ".html";
    try {
      await navigator.clipboard.writeText(url);
      linkCopied = true;
      setTimeout(() => (linkCopied = false), 2000);
    } catch {
      const input = document.createElement("input");
      input.value = url;
      document.body.appendChild(input);
      input.select();
      document.execCommand("copy");
      document.body.removeChild(input);
      linkCopied = true;
      setTimeout(() => (linkCopied = false), 2000);
    }
  };
</script>

<main class="flex-1 pr-4 pl-4 pt-4">
  <div class="max-w-2xl mx-auto">
    <div class="card p-8 text-center">
      <img
        class="w-24 h-24 rounded-full object-cover border-4 border-[#6b48ff] mx-auto mb-4"
        src={data.avatarUrl || "/avatars/" + data.id}
        alt={data.name}
      />
      <h1 class="text-2xl font-bold mb-1">{data.name}</h1>
      {#if data.title}
        <p class="text-gray-500 mb-6">{data.title}</p>
      {/if}

      <div class="flex flex-col sm:flex-row gap-3 justify-center">
        <a
          href={"/portfolios/" + data.id + ".html"}
          target="_blank"
          class="btn inline-flex items-center justify-center"
        >
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"></path>
          </svg>
          View Portfolio
        </a>
        <button class="btn-secondary inline-flex items-center justify-center" onclick={copyPortfolioLink}>
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3"></path>
          </svg>
          {linkCopied ? "Copied!" : "Copy Portfolio Link"}
        </button>
      </div>
    </div>
  </div>
</main>
