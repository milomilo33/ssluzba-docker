<template>
    <b-container>
        <h1>Registracija profesora</h1>
        <b-form>
            <b-form-input id="jmbg"
                        name="jmbg"
                        placeholder="JMBG"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="jmbg">
            </b-form-input>
            <br>
            <b-form-input id="ime"
                        name="ime"
                        placeholder="Ime"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="ime">
            </b-form-input>
            <br>
            <b-form-input id="prezime"
                        name="prezime"
                        placeholder="Prezime"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="prezime">
            </b-form-input>
            <br>
            <b-form-file
                v-model="image_url"
                :state="Boolean(image_url)"
                placeholder="Izaberi sliku ili je prevuci ovde..."
                drop-placeholder="Prevuci sliku ovde..."
                accept=".jpg, .png"
            ></b-form-file>
            <br>
            <br>
            <b-button @click="onSubmit" class="mb-2 mr-sm-2 mb-sm-0">Registruj</b-button>
        </b-form>

        <b-modal ref="error-modal" hide-footer title="Error">
            <div class="d-block text-center">
                <p>{{ this.errorMessage }}</p>
            </div>
            <b-button class="mt-3" variant="outline-danger" block @click="hideErrorModal">Close</b-button>
        </b-modal>
        
        <b-modal ref="success-modal" hide-footer title="Success">
            <div class="d-block text-center">
                <p>Uspešno registrovan profesor!</p>
            </div>
            <b-button class="mt-3" variant="outline-success" block @click="hideSuccessModal">Close</b-button>
        </b-modal>

    </b-container>
</template>

<script>

    export default {
        data() {
            return {
                jmbg: '',
                ime: '',
                prezime: '',
                image_url: null,
                errorMessage: '',
            }
        },
        methods: {
            onSubmit() {
                this.errorMessage = ""
                if (this.ime === "") {
                    this.errorMessage = "Ime ne može biti prazno";
                    this.showErrorModal();
                }
                if (this.prezime === "") {
                    this.errorMessage = "Prezime ne može biti prazno";
                    this.showErrorModal();
                }
                if (!this.jmbg) {
                    this.errorMessage = "JMBG ne može biti prazan";
                    this.showErrorModal();
                }

                if (this.errorMessage !== ""){
                    return
                }

                let formData = new FormData();
                formData.append("image_url", this.image_url, this.jmbg + '.' + this.image_url.name.split(".").pop());
                formData.append("ime", this.ime);
                formData.append("prezime", this.prezime);
                formData.append("jmbg", this.jmbg);

                this.axios.post(`http://localhost:8000/ftn/profesor/api`, formData, {
                        headers: {
                            'Content-Type': 'multipart/form-data',
                        },         maxContentLength: 100000000,
        maxBodyLength: 1000000000
                    })
                .then(() => {
                    this.showSuccessModal();
                })
                .catch(error => {
                    console.log(error);
                    this.errorMessage = "Neuspešna registracija!";
                    this.showErrorModal();
                });
            },

            hideErrorModal() {
                this.$refs['error-modal'].hide()
            },

            hideSuccessModal() {
                this.$refs['success-modal'].hide()
            },

            showErrorModal() {
                this.$refs['error-modal'].show()
            },

            showSuccessModal() {
                this.$refs['success-modal'].show()
            },
        }
    }
</script>
