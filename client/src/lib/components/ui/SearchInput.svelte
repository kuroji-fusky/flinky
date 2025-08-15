<script lang="ts">
  import { onMount, type Snippet } from "svelte";
  import { ChevronDownIcon, SearchIcon, XIcon } from "../icons";

  interface Props {
    defaultInputValue?: unknown[];
    bottomLayer?: Snippet;
    isDismissedProtip?: boolean;
  }

  const {
    defaultInputValue,
    isDismissedProtip = false,
    bottomLayer,
  }: Props = $props();

  // Used for advanced search
  const parsedInputValue = $state();

  let isInputFocused = $state(false);
  let inputHasContents = $state(false);

  let inputBox: HTMLInputElement;

  const focusInputOnClear = () => {
    inputHasContents = false;
    inputBox.value = "";
    inputBox.focus();
  };

  onMount(() => {
    const { signal, abort: inputEventAbort } = new AbortController();

    const contentHandler = (e: Event) => {
      const textInput = (e.target as HTMLInputElement).value;

      if (textInput !== "") {
        inputHasContents = true;
        return;
      }

      inputHasContents = false;
      return;
    };

    inputBox.addEventListener("input", contentHandler, { signal });

    return () => {
      inputEventAbort();
    };
  });
</script>

<div class="flex relative rounded-sm border">
  <div
    class="pointer-events-none absolute inset-y-0 left-0 grid place-items-center-safe px-2.5"
  >
    <SearchIcon />
  </div>
  <div id="scroller-input" class="flex-1">
    <input
      bind:this={inputBox}
      role="searchbox"
      placeholder="Search"
      type="text"
      name="root-search"
      class="w-full pl-9 pr-3 bg-inherit text-sm py-1.5 border-none"
    />
  </div>
  {#if inputHasContents}
    <button class="px-2 py-2 cursor-pointer" onclick={focusInputOnClear}>
      <XIcon />
    </button>
  {/if}
  <button class="py-1.5 px-2 cursor-pointer">
    <ChevronDownIcon />
  </button>
  <button class="py-1.5 px-3 bg-blue-500 cursor-pointer">
    <SearchIcon />
  </button>
</div>
<div id="bottom-layer-container" class="relative">
  {#if !isDismissedProtip}
    <div class="opacity-75 mt-2">
      <strong>Protip:</strong> <span>protip</span>
    </div>
  {/if}
  <div id="bottom-layer">
    {@render bottomLayer?.()}
  </div>
</div>
