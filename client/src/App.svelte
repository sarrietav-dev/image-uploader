<script lang="ts">
  import Card from "./components/Card.svelte";
  import axios from "axios";

  const onSubmit = (event: SubmitEvent) => {
    event.preventDefault();
    const formData = new FormData(event.target as HTMLFormElement);
    axios.post("http://localhost:8080/api/image", formData, {
      onUploadProgress: (progressEvent) => {
        console.log(progressEvent.loaded / progressEvent.total);
      },
    });
  };
</script>

<Card {onSubmit} />
