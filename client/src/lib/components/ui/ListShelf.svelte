<script>
  import { ChevronDownIcon, LayoutGridIcon, FilterIcon, Settings2Icon } from "../icons";
</script>

<div id="list-view-shelf" class="flex gap-x-3">
  <button
    class="md:hidden inline-flex items-center gap-x-1.5 px-2 py-2 cursor-pointer"
  >
    <span>Options</span>
    <Settings2Icon />
  </button>
  <div class="hidden md:flex gap-x-1">
    <button
      class="inline-flex items-center gap-x-1.5 px-2.5 py-1.5 border rounded-md cursor-pointer"
    >
      <FilterIcon width="1.4em" height="1.4em" />
      <span>Sort</span>
      <ChevronDownIcon width="1em" height="1em" />
    </button>
    <button
      class="inline-flex items-center gap-x-1.5 px-2.5 py-1.5 border rounded-md cursor-pointer"
    >
      <LayoutGridIcon width="1.4em" height="1.4em" />
      <span>View</span>
      <ChevronDownIcon width="1em" height="1em" />
    </button>
  </div>
</div>
