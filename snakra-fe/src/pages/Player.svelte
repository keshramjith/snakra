<script lang="ts">
  import { onMount } from "svelte";
  import NotFound from "./NotFound.svelte";

  const convertURIToBinary = (dataURI) => {
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
    const setAudio = (binary) => {
        let blob = new Blob([binary], { type: "audio/ogg; codecs=opus" })
        let blobURL = URL.createObjectURL(blob)
        return blobURL
    }
    export let id;
    const getVn = async (id: string) => {
        const response = await fetch (`http://localhost:3001/api/v1/vn/${id}`)
        const json = await response.json()
        return json.audio
    }
</script>

{#await getVn(id) then data }
    <audio src="{setAudio(convertURIToBinary(data))}" controls/>
{/await}
<p>id: {id}</p>