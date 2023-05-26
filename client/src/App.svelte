<script lang="ts">
  import axios from "axios";
  import ImageUploader from "./components/ImageUploader.svelte";
  import UploadedImageCard from "./components/UploadedImageCard.svelte";

  let state: "initial" | "loading" | "uploaded" = "initial";
  let imgId: string;

  const onSubmit = async (event: SubmitEvent) => {
    event.preventDefault();
    const formData = new FormData(event.target as HTMLFormElement);
    state = "loading";
    const response = await axios.post<{ imgId: string }>(
      "http://localhost:8080/api/image",
      formData,
      {
        onUploadProgress: (progressEvent) => {
          console.log(progressEvent.loaded / progressEvent.total);
        },
      }
    );

    state = "uploaded";
    imgId = response.data.imgId;
  };
</script>

{#if state === "initial"}
  <ImageUploader on:submit={onSubmit} />
{:else if state === "loading"}
  <p>Loading...</p>
{:else if state === "uploaded"}
  <UploadedImageCard {imgId} />
{/if}
