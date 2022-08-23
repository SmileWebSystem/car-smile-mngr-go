package soapRepo

//
//
//
var templateRequest = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://www.segurosbolivar.com/simon/fasecolda/1.0">
   <soapenv:Header/>
   <soapenv:Body>
      <ns:ConsultaSisaRequest>
         <DataHeader>
            <modulo>2</modulo>
            <proceso>2000</proceso>
            <subProceso>240</subProceso>
            <codCia>3</codCia>
            <codSecc>1</codSecc>
            <codProducto>250</codProducto>
            <!--Optional:-->
            <subProducto>1</subProducto>
            <codUrs>51938035</codUrs>
            <entidadColocadora>0</entidadColocadora>
            <!--Optional:-->
            <canal>1</canal>
            <sistemaOrigen>104</sistemaOrigen>
            <pais>1</pais>
            <direccionIp>127.0.0.1</direccionIp>
            <versionServicio>1.0</versionServicio>
            <!--Optional:-->
            
         </DataHeader>
         <!--Optional:-->
         <placa>{{.LicensePlate}}</placa>
         <!--Optional:-->         
      </ns:ConsultaSisaRequest>
   </soapenv:Body>
</soapenv:Envelope>`

//
//
//
type RequestParam struct {
	LicensePlate string
}

//
//
//
func populateRequest(licensePlate string) *RequestParam {
	req := RequestParam{}
	req.LicensePlate = licensePlate
	return &req
}
