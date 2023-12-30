<script lang="ts">
  import { page } from '$app/stores';
  import { goto } from "$app/navigation";

  const convertURIToBinary = (dataURI: string) => {
      let BASE64_MARKER = ';base64,';
      let base64Index = dataURI.indexOf(BASE64_MARKER) + BASE64_MARKER.length;
      let base64 = dataURI.substring(base64Index);
      let raw = window.atob(base64);
      let rawLength = raw.length;
      let arr = new Uint8Array(new ArrayBuffer(rawLength));
    
      for (let i = 0; i < rawLength; i++) {
          arr[i] = raw.charCodeAt(i);
      }
      return arr;
  }

  const getVn = async (id: string) => {
      const response = await fetch (`${import.meta.env.VITE_API_URL}/api/v1/vn/${id}`)
      if (response.status == 200) {
          const json = await response.json()
          return json.audio
      }
      if (response.status == 404) {
          goto("/notfound", { replaceState: true })
      }
  }

  const setAudio = (binary: Uint8Array) => {
      let blob = new Blob([binary], { type: "audio/ogg; codecs=opus" })
      let blobURL = URL.createObjectURL(blob)
      return blobURL
  }

  export let id: string

  id = $page.params.id;
  console.log(id)
</script>
  
{#await getVn(id) then data }
    <audio src="{setAudio(convertURIToBinary(data))}" controls/>
{/await}