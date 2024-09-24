'use client';

import { useCallback, useEffect, useState } from "react";

export default function Home() {

  const [haveMicPerm, setHaveMicPerm] = useState(false);
  const [audioRecorder, setAudioRecorder] = useState<MediaRecorder>();
  const [audioSrc, setAudioSrc] = useState('');
  const [recording, setRecording] = useState(false);

  const [chunks, setChunks] = useState<Blob[]>([]);

  const onDataAvailable = useCallback((ev: BlobEvent) => {
    const blob = new Blob([...chunks, ev.data], { type: 'audio/webm;codecs=opus' });
    const audioURL = window.URL.createObjectURL(blob);
    setAudioSrc(audioURL);
  }, [chunks])

  useEffect(() => {
    navigator.mediaDevices.getUserMedia({ audio: true }).then(stream => {
      const recorder = new MediaRecorder(stream);
      recorder.ondataavailable = onDataAvailable;
      recorder.onstop = () => {
        setChunks([])
      };

      setAudioRecorder(recorder);
      setHaveMicPerm(true);
    })
  }, [onDataAvailable, chunks])

  const recordButtonHandler = useCallback((e: React.MouseEvent) => {
    e.preventDefault()
    setRecording(true);
    if (audioRecorder)
      audioRecorder.start()
  }, [audioRecorder])

  const stopButtonHandler = useCallback((e: React.MouseEvent) => {
    e.preventDefault();
    setRecording(false)
    if (audioRecorder)
      audioRecorder.stop()
  }, [audioRecorder])

  return <>
    {!haveMicPerm ? <p>Need microphone access to continue</p> : (!recording ? 
      <button className="bg-blue-500 rounded" onClick={recordButtonHandler}>Record</button>
    : <button className="bg-red-500 rounded" onClick={stopButtonHandler}>Stop</button>)
    }
    {audioSrc ? <audio src={audioSrc} controls /> : <></>}
  </>;
}
