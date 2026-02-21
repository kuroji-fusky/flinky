<script lang="ts">
  import {
    ChevronRightIcon,
    ArrowRightIcon,
    RandomIcon,
    ChevronDownIcon,
  } from "../icons";

  const CATEGORIES = [
    "species",
    "franchise",
    "origin",
    "chronology",
    "medium",
    "target audience",
  ];

  let isExpanded = $state(false);
  let navContainer: HTMLDivElement;

  const handleContainerLeave = () => {
    if (isExpanded) {
      isExpanded = false;

      window.removeEventListener("keydown", handleKeyEvent);
      return;
    }
  };

  const handleKeyEvent = (e: KeyboardEvent) => {
    if (isExpanded && e.key === "Escape") {
      isExpanded = false;

      window.removeEventListener("keydown", handleKeyEvent);
      return;
    }
  };

  $effect(() => {
    if (!isExpanded) return;

    window.addEventListener("keydown", handleKeyEvent);
  });
</script>

<svelte:window onblur={() => (isExpanded = false)} />

<button
  onmouseenter={() => (isExpanded = true)}
  class="group px-2 h-full py-1.5 rounded-md cursor-pointer hidden lg:inline-flex items-center gap-x-1"
>
  <span>Browse</span>
  <ChevronDownIcon
    width="1em"
    height="1em"
    class="opacity-65 transition-transform group-hover:translate-y-0.5 group-hover:opacity-100"
  />
</button>

<div
  role="navigation"
  bind:this={navContainer}
  onmouseleave={handleContainerLeave}
  class={[
    "fixed inset-x-0 top-(--navbar-size) bg-black",
    !isExpanded ? "hidden" : "",
  ]}
>
  <div
    class="px-6 pt-2 rounded-md grid grid-cols-[auto_1fr] gap-4 ml-17 max-w-screen-2xl"
  >
    <nav class="grid gap-y-0.5 w-64 h-fit">
      <div class="py-2 px-2.5 opacity-50">Categories</div>
      {#each CATEGORIES as a}
        <button
          class="group inline-flex items-center-safe justify-between py-2 px-2.5 rounded-md hover:bg-neutral-500/10 dark:hover:bg-neutral-500/20 cursor-pointer"
        >
          <span>{`By ${a}`}</span>
          <ChevronRightIcon
            width="1.1em"
            height="1.1em"
            class="-translate-x-0.5 opacity-60 group-hover:translate-x-0 group-hover:opacity-100 transition-transform"
          />
        </button>
      {/each}
      <a
        href="/categories"
        class="group inline-flex items-center-safe justify-between py-2 px-2.5 rounded-md hover:bg-purple-500/10 dark:hover:bg-purple-500/20 text-purple-500"
        ><span>View all categories</span>
        <ArrowRightIcon
          width="1.1em"
          height="1.1em"
          class="-translate-x-0.5 opacity-60 group-hover:translate-x-0 group-hover:opacity-100 transition-transform"
        />
      </a>
      <hr class="my-1 mx-2 opacity-40" />
      <button
        class="group inline-flex items-center-safe gap-x-2 py-2 px-2.5 rounded-md hover:bg-pink-500/10 dark:hover:bg-pink-500/20 text-pink-500 cursor-pointer"
      >
        <RandomIcon width="1.25em" height="1.25em" />
        <span>Random</span>
      </button>
    </nav>
    <section
      data-dynamic-list-item=""
      class="grid grid-rows-[auto_1fr] gap-y-2"
      aria-label="Browse by species"
    >
      <h2 class="font-bold text-xl leading-relaxed">Species</h2>
      <div class="grid grid-cols-3 gap-2 h-full" role="listbox">
        {#each [...Array(12)] as _}
          <div
            role="listitem"
            class="flex items-center gap-x-2"
            aria-labelledby="sub-heading-<random-hash>"
            aria-describedby="sub-subtitle-<random-hash>"
          >
            <a href="/" aria-label="View item image">
              <div class="bg-red-200 size-18 rounded-md row-span-2"></div>
            </a>
            <div class="space-y-0.5">
              <h3
                class="text-xl font-bold h-fit"
                id="sub-heading-<random-hash>"
              >
                <a href="/">Title</a>
              </h3>
              <span
                class="opacity-75 cursor-default"
                id="sub-subtitle-<random-hash>"
              >
                Subheading
              </span>
            </div>
          </div>
        {/each}
      </div>
    </section>
    <div class="col-span-2 pt-2 pb-4 px-2">
      <p>
        <span class="opacity-75">Can't find what you're looking for? Try </span>
        <a href="/search" class="hover:underline text-blue-400"
          >Advanced search</a
        >
      </p>
    </div>
  </div>
</div>
