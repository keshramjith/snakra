package com.snakra.audioservice;

import org.springframework.web.bind.annotation.*;

import java.io.FileOutputStream;
import java.util.Base64;

@CrossOrigin
@RestController
public class AudioController {
    @RequestMapping(value = "/api", method = RequestMethod.POST)
    public void PostAudio(@RequestBody String base64Audio) throws Exception {
        Base64.Decoder decoder = Base64.getDecoder();
        byte[] decodedByte = decoder.decode(base64Audio.split(",")[1]);
        FileOutputStream fos = new FileOutputStream("MyAudio.webm");
        fos.write(decodedByte);
        fos.close();
    }
}
