package entity

import "tracerstudy-siak-service/pb"

type MhsBiodataApi struct {
	NIM        string `json:"nim"`
	NAMA       string `json:"nama"`
	IPK        string `json:"ipk"`
	HP         string `json:"hp"`
	EMAIL      string `json:"email"`
	KODEPSTD   string `json:"kodepstd"`
	JENJANG    string `json:"jenjang"`
	PRODI      string `json:"prodi"`
	NAMAPST    string `json:"namapst"`
	KODEPST    string `json:"kodepst"`
	KODEFAK    string `json:"kodefak"`
	NAMAFAK    string `json:"namafak"`
	JLRMASUK   string `json:"jlrmasuk"`
	THNMASUK   string `json:"thnmasuk"`
	LAMASTD    string `json:"lamastd"`
	TGLSIDANG  string `json:"tglsidang"`
	KODEJK     string `json:"kodejk"`
	KODESTATUS string `json:"kodestatus"`
}

func NewMhsBiodataApi(nim, nama, ipk, hp, email, kodepstd, jenjang, prodi, namapst, kodepst, kodefak, namafak, jlrmasuk, thnmasuk, lamastd, tglsidang, kodejk, kodestatus string) *MhsBiodataApi {
	return &MhsBiodataApi{
		NIM:        nim,
		NAMA:       nama,
		IPK:        ipk,
		HP:         hp,
		EMAIL:      email,
		KODEPSTD:   kodepstd,
		JENJANG:    jenjang,
		PRODI:      prodi,
		NAMAPST:    namapst,
		KODEPST:    kodepst,
		KODEFAK:    kodefak,
		NAMAFAK:    namafak,
		JLRMASUK:   jlrmasuk,
		THNMASUK:   thnmasuk,
		LAMASTD:    lamastd,
		TGLSIDANG:  tglsidang,
		KODEJK:     kodejk,
		KODESTATUS: kodestatus,
	}
}

func ConvertEntityToProto(m *MhsBiodataApi) *pb.MhsBiodataApi {
	return &pb.MhsBiodataApi{
		NIM:        m.NIM,
		NAMA:       m.NAMA,
		IPK:        m.IPK,
		HP:         m.HP,
		EMAIL:      m.EMAIL,
		KODEPSTD:   m.KODEPSTD,
		JENJANG:    m.JENJANG,
		PRODI:      m.PRODI,
		NAMAPST:    m.NAMAPST,
		KODEPST:    m.KODEPST,
		KODEFAK:    m.KODEFAK,
		JLRMASUK:   m.JLRMASUK,
		THNMASUK:   m.THNMASUK,
		LAMASTD:    m.LAMASTD,
		TGLSIDANG:  m.TGLSIDANG,
		KODEJK:     m.KODEJK,
		KODESTATUS: m.KODESTATUS,
	}
}
