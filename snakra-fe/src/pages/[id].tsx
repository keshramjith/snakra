import { Blob } from "buffer"
import { Router, useRouter } from "next/router"
import "react"
import { useEffect, useState } from "react"

const ListenPage = () => {
    const router = useRouter()
    const { id } = router.query
    const [haveAudio, setHaveAudio] = useState(false)
    useEffect(() => {
        const getBlob = async () => {
            const resp = await fetch(`http://localhost:3001/${id}`)
            const blobR = await resp.blob()
            const addAudioElement = (blob) => {
                const url = URL.createObjectURL(blob);
                const audio = document.createElement("audio");
                audio.src = url;
                audio.controls = true;
                document.body.appendChild(audio);
              };
            addAudioElement(blobR)
        }
        if (!router.isReady) return
        getBlob()
    }, [router.isReady, router.query.id])

    return (
        <>
        </>
    )
}

export default ListenPage