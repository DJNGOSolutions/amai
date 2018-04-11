import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
// We have to import the components that we are gonna use.
import { HomeComponent }        from './home/home.component';
import { CrudComponent }        from './crud/crud.component';

// Pretty self explicatory we specify a route and it's component.
const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'crud', component: CrudComponent }
];

@NgModule({
    imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
