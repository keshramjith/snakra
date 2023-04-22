import { Blob } from "buffer"
import { useRouter } from "next/router"
import { useEffect, useRef } from "react"

const ListenPage = () => {
    const router = useRouter()
    const { id } = router.query
    const audioElement = useRef(null)
    useEffect(() => {
        if (!router.isReady) {
            console.log("nothing should happen")
            return
        }
        const getBlob = async () => {
            const resp = await fetch(`http://localhost:3333/${id}`)
            const blobR = await resp.blob()
            const addAudioElement = (blob: Blob) => {
                const url = URL.createObjectURL(blob);
                const { current } = audioElement
                if (current) {
                    current.src = url
                    current.controls = true
                }
              };
            addAudioElement(blobR)
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