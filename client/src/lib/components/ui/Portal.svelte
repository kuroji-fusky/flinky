<script lang="ts">
  import type { WithSnippet } from "../../utils";

  interface Props {
    to?: string;
  }

  const moveDom = (node: Element, inject: string) => {
    const target = document.querySelector(inject)!;
    target.appendChild(node);

    return {
      destroy() {
        node.remove();
      },
    };
  };

  const { to = "body", children }: WithSnippet<Props> = $props();
</script>

<div id="portal" class="contents empty:hidden" use:moveDom={to}>
  {@render children?.()}
</div>
