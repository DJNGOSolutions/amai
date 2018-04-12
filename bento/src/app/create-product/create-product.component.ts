import { Component, OnInit } from '@angular/core';
import { AmaiService } from '../amai.service';
import { Products } from '../Products';
import { EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-create-product',
  templateUrl: './create-product.component.html',
  styleUrls: ['./create-product.component.css']
})
export class CreateProductComponent implements OnInit {
    product : Products = {
        code: 'nil',
        price: 0
    };
    @Output()
    created = new EventEmitter<string>();

    createProduct( code: string, price : number ){
        console.log("on create");
        this.service.createProduct(code,price);
        this.product = {
            code: '',
            price: 0
        };
        //this.created.emit('created');
    }

    constructor( private service: AmaiService) { }

    ngOnInit() {
    }

}
