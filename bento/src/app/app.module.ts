import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { CrudComponent } from './crud/crud.component';
import { AppRoutingModule } from './app-routing.module';
import { AmaiService } from './amai.service';
import { CreateProductComponent } from './create-product/create-product.component';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    CrudComponent,
    CreateProductComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [AmaiService],
  bootstrap: [AppComponent]
})
export class AppModule { }