<script lang="ts">
  let { header, detail }: { header: string; detail: string } = $props();

  let isOpen = $state(true);

  const cancelModel = () => {
    isOpen = false;
  };
</script>

{#if isOpen}
  <!-- Portfolio Generation Modal -->
  <div id="portfolioModal" class="modal hidden">
    <div class="bg-white rounded-lg w-1/2 max-w-2xl shadow-lg text-center">
      <div class="p-6">
        <!-- svg dấu tick
        <svg
          class="w-16 h-16 text-[#6b48ff] mx-auto mb-4"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
          ></path>
        </svg> -->

        <!-- svg dấu chấm hỏi -->
        <svg
          class="w-16 h-16 text-[#6b48ff] mx-auto mb-4"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          ></path>
        </svg>
        <h3 class="text-xl font-semibold mb-2">
          Are you sure about creating these tasks?
        </h3>
        <p class="text-gray-600 mb-4">
          4 new tasks will be created, and you can edit, delete these tasks in
          the future
        </p>
        <!-- <div class="bg-gray-100 p-4 rounded mb-4 text-left">
          <p class="text-sm font-mono break-all">
            https://skillforge.example.com/portfolio/johndoe
          </p>
        </div> -->
        <div class="flex space-x-2 justify-center">
          <button
            id="closePortfolioBtn"
            class="btn-secondary"
            onclick={cancelModel}>Cancel</button
          >
          <button class="btn flex items-center" onclick={cancelModel}>
            Yes
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}

<style>
  .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 50;
  }
</style>
