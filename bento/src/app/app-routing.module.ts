import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent }        from './home/home.component';
import { CrudComponent }        from './crud/crud.component';

const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'crud', component: CrudComponent }
];

@NgModule({
    imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
