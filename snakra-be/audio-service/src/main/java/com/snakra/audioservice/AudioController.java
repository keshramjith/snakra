package com.snakra.audioservice;

import com.snakra.audioservice.models.GetResponse;
import com.snakra.audioservice.models.PostResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.util.Base64Utils;
import org.springframework.web.bind.annotation.*;
import software.amazon.awssdk.core.ResponseInputStream;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.GetObjectRequest;
import software.amazon.awssdk.services.s3.model.PutObjectRequest;
import java.util.Base64;
import java.util.UUID;

@CrossOrigin
@RestController
public class AudioController {
    private final S3Client s3Client;

    public AudioController() {
        s3Client = DepedencyFactory.s3Client();
    }

    @RequestMapping(value = "/api", method = RequestMethod.POST)
    public ResponseEntity<PostResponse> PostAudio(@RequestBody String base64Audio) {
        Base64.Decoder decoder = Base64.getDecoder();
        byte[] decodedByte = decoder.decode(base64Audio.split(",")[1]);
        String uuid = UUID.randomUUID().toString();
        s3Client.putObject(PutObjectRequest.builder()
                .bucket("snakra-test")
                .key(uuid)
                .build(),
                software.amazon.awssdk.core.sync.RequestBody.fromBytes(decodedByte));
        PostResponse resp = new PostResponse();
        resp.id = uuid;
        return new ResponseEntity<>(resp, HttpStatus.OK);
    }

    @RequestMapping(value = "/api/{uuid}", method = RequestMethod.GET)
    public ResponseEntity<GetResponse> GetAudio(@PathVariable String uuid) throws Exception {
        ResponseInputStream s3response = s3Client.getObject(GetObjectRequest.builder().bucket("snakra-test").key(uuid).build());
        byte[] audio = s3response.readAllBytes();
        Base64.Encoder encoder = Base64.getEncoder();
        GetResponse resp = new GetResponse();
        resp.audio = encoder.encodeToString(audio);
        return new ResponseEntity<>(resp, HttpStatus.OK);
    }
}
