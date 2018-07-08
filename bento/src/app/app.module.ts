import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { CrudComponent } from './crud/crud.component';
import { AppRoutingModule } from './app-routing.module';
import { AmaiService } from './amai.service';
import { PanelComponent } from './panel/panel.component';
import { MatGridListModule, MatCardModule, MatMenuModule, MatIconModule,
    MatButtonModule, MatToolbarModule, MatSidenavModule, MatListModule,
    MatTableModule, MatPaginatorModule, MatSortModule, MatInputModule, 
    MatSelectModule,MatOptionModule,MatNativeDateModule } from '@angular/material';
import {MatDatepickerModule, } from '@angular/material/datepicker';
import {MatFormFieldModule } from '@angular/material/form-field';
import { LayoutModule } from '@angular/cdk/layout';
import { MainNavComponent } from './main-nav/main-nav.component';
import { UsersTableComponent } from './users-table/users-table.component';
import { CreateUserComponent } from './create-user/create-user.component';
import { LoginComponent } from './login/login.component';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    CrudComponent,
    PanelComponent,
    MainNavComponent,
    UsersTableComponent,
    CreateUserComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    AppRoutingModule,
    HttpClientModule,
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule,
    LayoutModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatOptionModule,
    MatDatepickerModule,
    MatNativeDateModule
  ],
  providers: [AmaiService],
  bootstrap: [AppComponent]
})
export class AppModule { }
