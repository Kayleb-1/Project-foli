<script lang="ts">
  import Grid from 'svelte-grid';
  import gridHelp from 'svelte-grid/build/helper/index.mjs';

  // Generate unique IDs for grid items
  const id = () => '_' + Math.random().toString(36).substr(2, 9);

  // Define grid items with custom data and properties
  let items = [
    {
      6: gridHelp.item({
        x: 0,
        y: 0,
        w: 2,
        h: 2,
        draggable: false, // Non-draggable item
        max: { w: 2, h: 2 },
      }),
      id: id(),
      data: { title: 'Fixed Item', color: 'bg-blue-200', content: 'This item is fixed.' },
    },
    {
      6: gridHelp.item({
        x: 2,
        y: 0,
        w: 2,
        h: 3,
        resizable: true, // Resizable item
        draggable: true,
      }),
      id: id(),
      data: { title: 'Resizable Item', color: 'bg-green-200', content: 'Drag and resize me!' },
    },
    {
      6: gridHelp.item({
        x: 4,
        y: 0,
        w: 2,
        h: 2,
        draggable: true, // Draggable only
        resizable: false,
      }),
      id: id(),
      data: { title: 'Draggable Item', color: 'bg-yellow-200', content: 'Drag me around!' },
    },
  ];

  // Define columns for responsive layout (6 columns at >=1200px)
  const cols = [[1200, 6]];
</script>

<!-- Grid component with gap, row height, and slot rendering -->
<Grid gap={[10, 5]} cols={cols} rowHeight={150} bind:items={items} let:item let:dataItem>
  <div class="{dataItem.data.color} border border-gray-300 rounded-lg p-4 h-full flex flex-col justify-between">
    <h3 class="font-bold">{dataItem.data.title}</h3>
    <p>{dataItem.data.content}</p>
    <span class="text-sm text-gray-600">ID: {dataItem.id}</span>
  </div>
</Grid>

<style>
  /* Additional styles if not using Tailwind */
  :global(.svlr-grid-item) {
    transition: all 0.2s ease;
  }
</style>