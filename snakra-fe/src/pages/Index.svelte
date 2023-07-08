<script lang="ts">
    let isRecording = false
    let micAccess = false
    let chunks = []
    let blob = null
    let recorder
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

    const upload = (blob: Blob) => {
        const reader = new FileReader()
        reader.readAsDataURL(blob)
        reader.onloadend = () => {
            const b64str = reader.result.split(',')[1]
            console.log("uploading: ", b64str)
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
            <button on:click={e => stopRecordingHandler(e)}>Stop</button>
        {:else}
            {#if blob}
                <audio src={ audioSrc } controls />
                <button on:click={e => shareButtonHandler(e)}>Make shareable</button>
                <button on:click={e => resetHandler(e)}>Reset</button>
            {:else}
                <button on:click={e => startRecordingHandler(e)}>Record</button>
            {/if}
        {/if}
    {:else}
        <p>Microphone access is needed.</p>
    {/if}
</div>