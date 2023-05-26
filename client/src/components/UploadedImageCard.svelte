<script lang="ts">
  import Button from "./Button.svelte";
  import Card from "./Card.svelte";
  import CheckIcon from "./CheckIcon.svelte";
  export let imgId: string;
  let url = `http://localhost:8080/api/image/${imgId}`;

  const copyToClipboard = () => {
    navigator.clipboard.writeText(url);
  };
</script>

<Card>
  <CheckIcon />
  <h1 class="header">Uploaded Successfully!</h1>
  <div class="img-container">
    <img src={url} alt="Uploaded media" />
  </div>
  <div class="link-container">
    <div class="link">
      <a href={url} target="_blank">
        {url}
      </a>
    </div>
    <div class="btn">
      <Button on:click={copyToClipboard}>Copy Link</Button>
    </div>
  </div>
</Card>

<style>
  .header {
    margin-top: 1rem;
  }

  .img-container {
    height: 200px;
    width: 75%;
    border-radius: 12px;
    margin: 1rem 0 2rem 0;
  }

  .img-container img {
    height: 100%;
    width: 100%;
    object-fit: cover;
    border-radius: 12px;
  }

  .link-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border: 1px solid #e0e0e0;
    border-radius: 8px;
    font-size: 8px;
    max-width: 75%;
    padding: 2px 8px;
    padding-right: 0;
  }

  .link-container a {
    color: #4f4f4f;
    text-decoration: none;
    white-space: nowrap;
  }

  .link-container .link {
    overflow: hidden;
    text-overflow: ellipsis;
    margin-right: auto;
  }

  .link-container .btn {
    width: 41%;
  }
</style>
