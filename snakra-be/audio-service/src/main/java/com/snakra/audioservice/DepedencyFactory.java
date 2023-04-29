package com.snakra.audioservice;

import software.amazon.awssdk.auth.credentials.ProfileCredentialsProvider;
import software.amazon.awssdk.services.s3.S3Client;

public class DepedencyFactory {

    private DepedencyFactory() {}

    public static S3Client s3Client() {
        return S3Client.builder()
                .credentialsProvider(ProfileCredentialsProvider.create("default"))
                .build();
    }
}
