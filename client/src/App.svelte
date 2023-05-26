<script lang="ts">
  import axios from "axios";
  import ImageUploader from "./components/ImageUploader.svelte";

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

<ImageUploader {onSubmit} />
