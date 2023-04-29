import { useRouter } from "next/router"
import { useEffect, useRef } from "react"

const ListenPage = () => {
    const router = useRouter()
    const { id } = router.query
    const audioElement = useRef(null)
    useEffect(() => {
        if (!router.isReady) {
            return
        }
        const getBlob = async () => {
            const resp = await fetch(`http://localhost:8080/api/${id}`)
            const { audio } = await resp.json()
            const blob = new Blob([Buffer.from(audio, "base64")], { type: "audio/webm;codecs=opus;base64"})
            const addAudioElement = (blob: Blob) => {
                const url = URL.createObjectURL(blob);
                const { current } = audioElement
                if (current) {
                    current.src = url
                    current.controls = true
                }
              };
            addAudioElement(blob)
        }
        getBlob()
    }, [router.isReady, router.query.id])

    return (
        <>
            <audio ref={audioElement} />
        </>
    )
}

export default ListenPage