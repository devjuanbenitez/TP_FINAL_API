package com.testing.inventario.entities.error;

import lombok.Getter;
import lombok.Setter;
@Getter
@Setter
public class ApiErrorResponse {

    private int code;
    private String message;
}
