import { Component, OnInit } from '@angular/core';
import { AmaiService } from '../amai.service';
import { User } from '../User';
import { EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-create-product',
  templateUrl: './create-product.component.html',
  styleUrls: ['./create-product.component.css']
})
export class CreateProductComponent implements OnInit {
    user : User = {
        NombreUsuario: 'nil',
        EdadUsuario: 0
    };
    @Output()
    created = new EventEmitter<string>();

    createProduct( name: string, age : number ){
        console.log("on create");
        this.service.createProduct(name,age);
        this.user = {
            NombreUsuario: '',
            EdadUsuario: 0
        };
        this.created.emit('created');
    }

    constructor( private service: AmaiService) { }

    ngOnInit() {
    }

}
