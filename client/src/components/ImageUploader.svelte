<script lang="ts">
  import Card from "./Card.svelte";

  import image from "../assets/image.svg";

  export let onSubmit: (event: SubmitEvent) => void;

  let fileInput: HTMLInputElement;
  let form: HTMLFormElement;

  const onChooseFile = () => {
    fileInput.click();
  };
</script>

<Card>
  <h1 class="card__header">Upload your file</h1>
  <p>File should be JPEG or PNG</p>
  <form
    bind:this={form}
    method="POST"
    enctype="multipart/form-data"
    on:submit={onSubmit}
  >
    <div class="card__img-drop">
      <img src={image} alt="Mountains" />
      <p>Drag & Drop your image here.</p>
    </div>
    <p>Or</p>
    <input
      on:change={() => {
        form.dispatchEvent(new Event("submit", { cancelable: true }));
      }}
      bind:this={fileInput}
      class="file-input"
      type="file"
      name="image"
    />
    <button type="button" on:click={onChooseFile} class="btn"
      >Choose a file</button
    >
  </form>
</Card>

<style>
  .card__header {
    font-size: 18px;
    color: #4f4f4f;
    letter-spacing: -0.035em;
    font-weight: 500;
    margin-bottom: 1rem;
  }

  .card__header + p {
    font-size: 10px;
    color: #828282;
  }

  .card__img-drop {
    border: 2px dashed #97bef4;
    border-radius: 10px;
    padding: 3rem 5rem;
    margin: 2rem 0;
    background-color: #f6f8fb;
  }

  .card__img-drop img {
    margin-bottom: 1rem;
  }

  .card__img-drop p {
    font-size: 12px;
    line-height: 18px;
    color: #bdbdbd;
    user-select: none;
  }

  .card__img-drop + p {
    font-size: 12px;
    line-height: 18px;
    color: #bdbdbd;
  }

  .file-input {
    display: none;
  }

  .btn {
    background-color: #2f80ed;
    border-radius: 8px;
    color: #fff;
    padding: 0.5rem 1rem;
    border: none;
    font-size: 12px;
    line-height: 18px;
    cursor: pointer;
    font-family: "Noto Sans", sans-serif;
  }

  .btn {
    margin-top: 1rem;
  }
</style>
