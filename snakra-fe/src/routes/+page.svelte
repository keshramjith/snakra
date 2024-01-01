<script lang="ts">
    import Icon from "@iconify/svelte"
    import { goto } from "$app/navigation"
    let isRecording = false
    let micAccess = false
    let chunks: any[] = []
    let blob = null
    let recorder: MediaRecorder
    let audioSrc

    navigator.mediaDevices.getUserMedia({ audio: true }).then(stream => {
        recorder = new MediaRecorder(stream)
        micAccess = true
        recorder.ondataavailable = (e) => {
            chunks.push(e.data);
        };
        recorder.onstop = e => {
            blob = new Blob(chunks, { type: "audio/ogg; codecs=opus" });
            chunks = []
            const audioURL = window.URL.createObjectURL(blob);
            audioSrc = audioURL;
        }
    })
    const startRecordingHandler = (e: MouseEvent) => {
        e.preventDefault()
        console.log("recording....")
        isRecording = true
        recorder.start()
    }
    const stopRecordingHandler = (e: MouseEvent) => {
        e.preventDefault()
        console.log("stopped recording")
        isRecording = false
        recorder.stop()
    }

    const upload = async (blob: Blob) => {
        const reader = new FileReader()
        reader.readAsDataURL(blob)
        reader.onloadend = async () => {
            const b64str = reader.result
            const reqBody = {
                vnString: b64str
            }
            const resp = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/vn`, {
                method: "POST",
                body: JSON.stringify(reqBody)
            })
            const json = await resp.json()
            goto(`/${json.url_short_form}`)
        }
    }

    const shareButtonHandler = (e: MouseEvent) => {
        e.preventDefault()
        upload(blob)
    }

    const resetHandler = (e: MouseEvent) => {
        e.preventDefault()
        blob = null
    }
</script>

<div>
    {#if micAccess}
        {#if isRecording}
            <div>
                <p>Recording</p>
                <Icon icon="ph:record-fill" width=35 color='red' />
            </div>
            <button on:click={e => stopRecordingHandler(e)}>Stop</button>
        {:else}
            {#if blob}
                <audio src={ audioSrc } controls />
                <div>
                    <button on:click={e => shareButtonHandler(e)}>Make shareable</button>
                    <button on:click={e => resetHandler(e)}>Reset</button>
                </div>
            {:else}
                <button on:click={e => startRecordingHandler(e)}>
                    <Icon icon="ph:record-fill" width=35 />
                </button>
            {/if}
        {/if}
    {:else}
        <p>Microphone access is needed.</p>
    {/if}
</div>