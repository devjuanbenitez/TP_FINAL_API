package com.testing.inventario.rest;

import com.testing.inventario.entities.dto.CategoriaDTO;
import com.testing.inventario.entities.dto.MarcaDTO;
import com.testing.inventario.entities.dto.ProductoDTO;
import com.testing.inventario.services.CategoriaService;
import com.testing.inventario.services.ProductoService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value ="${base.api}/productos",produces = MediaType.APPLICATION_JSON_VALUE)
@RequiredArgsConstructor
public class ProductoResources {

    private final ProductoService productoService;

    @PostMapping
    public ResponseEntity<String> createProducto(
            @RequestBody final ProductoDTO productoDTO) {
        productoService.create(productoDTO);
        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @GetMapping("/{id}")
    public ResponseEntity<ProductoDTO> getProductoById(
            @PathVariable(name = "id") final Long id) {
        return ResponseEntity.ok(productoService.getById(id));
    }

    @GetMapping()
    public List<ProductoDTO> getAllProducto() {
        return productoService.getList();
    }

    @PutMapping("/{id}")
    public ResponseEntity<Void> updateProducto(
            @PathVariable(name = "id") final Long id,
            @RequestBody ProductoDTO productoDTO) {
        productoService.update(id, productoDTO);
        return ResponseEntity.ok().build();
    }
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteProducto(
            @PathVariable(name = "id") final Long id) {
        productoService.delete(id);
        return ResponseEntity.noContent().build();
    }

}
