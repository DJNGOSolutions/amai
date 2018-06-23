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
        UserName: 'nil',
        UserAge: 0,
        UserEmail: ''
    };
    @Output()
    created = new EventEmitter<string>();

    createProduct( name: string, age : number, userEmail : string ){
        console.log("on create");
        this.service.createProduct(name,age,userEmail);
        this.user = {
            UserName: '',
            UserAge: 0,
            UserEmail: ''
        };
        this.created.emit('created');
    }

    constructor( private service: AmaiService) { }

    ngOnInit() {
    }

}
