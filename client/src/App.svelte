<script lang="ts">
  import Card from "./components/Card.svelte";

  const onSubmit = (event: SubmitEvent) => {
    event.preventDefault();
    const formData = new FormData(event.target as HTMLFormElement);
    const file = formData.get("image") as File;
    console.log(file);
  };
</script>

<Card {onSubmit} />
