import { Component, OnInit } from '@angular/core';
import { AmaiService } from '../amai.service';
import { Products } from '../Products';

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

    createProduct( code: string, price : number ){
        console.log("on create");
        this.service.createProduct(code,price);
    }

    constructor( private service: AmaiService) { }

    ngOnInit() {
    }

}
