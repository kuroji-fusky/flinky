<!-- shared -->
<script>
  import {
    MenuIcon,
    ChevronDownIcon,
    Settings2Icon,
    SearchIcon,
  } from "../icons";
  import BrowseMenu from "./BrowseMenu.svelte";
</script>

<nav class="fixed top-0 inset-x-0 z-100 bg-white dark:bg-black">
  <div
    class="h-[calc(var(--navbar-size)*1)] px-6 flex items-center gap-x-3.5 w-full"
  >
    <button class="px-2 py-2 rounded-md cursor-pointer block lg:hidden">
      <MenuIcon />
    </button>
    <div class="inline-flex items-center gap-x-1.5">
      <a href="/">Flinky</a>
      <button
        title="Expand"
        class="group p-1.5 hover:bg-neutral-500/50 rounded-sm"
      >
        <ChevronDownIcon
          width="1em"
          height="1em"
          class="opacity-65 transition-transform group-hover:opacity-100"
        />
      </button>
    </div>
    <span class="h-4 border-l opacity-40"></span>
    <div class="inline-flex gap-x-1">
      <BrowseMenu />

      <button
        class="group px-2 h-full py-1.5 rounded-md cursor-pointer hidden lg:inline-flex items-center gap-x-1"
      >
        <span>Explore</span>
        <ChevronDownIcon
          width="1em"
          height="1em"
          class="opacity-65 transition-transform group-hover:translate-y-0.5 group-hover:opacity-100"
        />
      </button>
      <button
        class="group px-2 h-full py-1.5 rounded-md cursor-pointer hidden lg:inline-flex items-center gap-x-1"
      >
        <span>Roadmap</span>
      </button>
      <button
        class="group px-2 h-full py-1.5 rounded-md cursor-pointer hidden lg:inline-flex items-center gap-x-1"
      >
        <span>Submissions</span>
      </button>
      <button
        class="group px-2 h-full py-1.5 rounded-md cursor-pointer hidden lg:inline-flex items-center gap-x-1"
      >
        <span>About</span>
      </button>
      <!--  -->
      <button
        id="search"
        aria-expanded="false"
        class="px-2.5 py-1.5 bg-blue-600/20 hover:bg-blue-600/60 rounded-md cursor-pointer inline-flex items-center gap-x-1"
        aria-label="Search or press control or command + slash for quick access"
      >
        <SearchIcon class="opacity-65" />
        <span>Search</span>
        <div>
          <span class="sr-only">Ctrl/Command + /</span>
          <kbd class="select-none text-xs border py-0.5 px-1 rounded-sm"
            >Ctrl</kbd
          >
          <kbd class="select-none text-xs border py-0.5 px-1 rounded-sm">/</kbd>
        </div>
      </button>
    </div>

    <div class="flex-1"></div>

    <div class="flex gap-x-1">
      <a href="/requests" class="px-3 py-1.5 border rounded-md cursor-pointer"
        >Request</a
      >
      <button
        class="px-2 py-2 border rounded-md cursor-pointer"
        aria-label="Site options and accessibility settings"
      >
        <Settings2Icon />
      </button>
    </div>
  </div>
</nav>
